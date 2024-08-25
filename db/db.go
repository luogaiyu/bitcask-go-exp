package db

import (
	"bitcask-go-exp/db/utils"
	"encoding/binary"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
20240825:
todo:
db实例剩下的方法:
	1. 获取全部的 Key
	2. 遍历所有的数据, 执行函数 Fun
	3. 执行merge
	4. Sync: 刷盘, 持久化到内存中
需要随着数据条数或者文件的大小切换文件
*/

// 注意体会go语言的 值传递和引用传递
type DB struct {
	index  *Btree        // 保存当前数据库的索引
	logPos *LogPos       // 保存当前数据库对应文件日志的地址, db需要维护这个pos来维护当前的db写入的位置
	fio    *FileIO       // 用来存放一些和文件交互的逻辑
	option *utils.Option // 用来保存一些数据内核中的配置项
}

// 初始化的方法
// merge的时候 才会新建索引的hint文件
// DataFile
func InitDB() *DB {
	// loadIndexFromFile() 从文件中重建索引
	db := &DB{
		index:  InitBTree(),
		logPos: InitLogPos(), // 这个地方应该是
		option: utils.InitTmpOptionForTest(),
	}
	db.fio = InitFileIO(db.option.DirPath)
	files, _ := os.ReadDir(db.option.DirPath)
	if len(files) == 0 {
		os.Create(filepath.Join(db.option.DirPath, "00000001.data"))
	} else {
		// 按照数据文件初始化变量
		db.loadIndexFromDataFile()
	}

	return db
}

func (db *DB) Put(key []byte, value []byte) {
	// 首先将数据保存到索引中, 再将数据写入到内存中, 为什么不将所有的数据写入到内存中, 如果全部写入到内存中可以存多少条数据?
	// todo 这里可以先判断这个数据在索引中是否提前存在
	db.index.Put(key, *db.logPos) // 这里[db.logPos] 需要使用值传递, 后面的修改和无关

	// lgPos LogPos, logRecord LogRecord
	lgRecord := &LogRecord{
		RType: LogRecordNormal, // 记录类型, 分为0和1
		// 20240825 修复bug: unsafe.sizeOf 指的是数组的指针的长度
		KeySize:   uint32(len(key)),
		ValueSize: uint32(len(value)),
		// 20240825 修复bug: unsafe.sizeOf 指的是数组的指针的长度
		Key:   key,
		Value: value,
	}
	err, n := db.fio.Write(*db.logPos, *lgRecord) // 这个地方应该每次都需要
	if err != nil {
		log.Fatal("db put log make error!")
	}
	// 写文件 -20240824 发现bug offset 导致读取数据位置不对
	db.logPos.Offset += uint64(n)
}

func (db *DB) Get(key []byte) ([]byte, error) {
	itm, err := db.index.Get(key)
	if err == nil {
		lgRecord := db.fio.Read(itm.(*Item).lgPos)
		return lgRecord.Value, nil
	} else {
		return nil, err
	}

}

// 不仅对当前索引中的数据进行删除, 并且写入一条日志
func (db *DB) Delete(key []byte) {
	db.index.Delete(key)
	lgRecord := &LogRecord{
		RType:     LogRecordNormal, // 记录类型, 分为0和1
		KeySize:   uint32(len(key)),
		ValueSize: uint32(0),
		Key:       key,
		Value:     nil,
	}
	err, n := db.fio.Write(*db.logPos, *lgRecord)
	if err != nil {
		log.Fatal("db delete log make error!")
	}
	db.logPos.Offset += uint64(n)
}

// 从数据文件中构建索引 || 还差构建一个删除的功能
// 并且需要写出多个文件来防止文件热点
func (db *DB) loadIndexFromDataFile() {
	dirEntrys, err := os.ReadDir(db.option.DirPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	var fileIds []string
	for _, dirEntry := range dirEntrys {
		if strings.HasSuffix(dirEntry.Name(), ".data") {
			fileId := strings.TrimSuffix(dirEntry.Name(), ".data")
			fileIds = append(fileIds, fileId)
		}
	}

	for i := 0; i < len(fileIds); i += 1 {
		// fileId 指的是data路径下的文件
		fileId := fileIds[i]
		file, err := os.OpenFile(filepath.Join(db.option.DirPath, fileId+".data"), os.O_RDONLY, 0666)// 这里需要改下, fileId 需要变成 00000001.data
		if err != nil {
		}

		lgPos := &LogPos{}
		cursor := int64(0)

		for {
			lgPos = &LogPos{
				Fid:    fileId,
				Offset: uint64(cursor),
			}
			// 一个个 record读取
			rtype_keySize_valueSize_value := make([]byte, 9)
			n, _ := file.ReadAt(rtype_keySize_valueSize_value, cursor)
			if n == 0 {
				break
			}
			keySize := binary.BigEndian.Uint32(rtype_keySize_valueSize_value[1:5])
			valueSize := binary.BigEndian.Uint32(rtype_keySize_valueSize_value[5:9])
			cursor += int64(n)
			key := make([]byte, keySize)
			n, _ = file.ReadAt(key, cursor)
			cursor += int64(n) + int64(keySize) + int64(valueSize)

			// 设置新的pos
			itm := &Item{
				key:   key,
				lgPos: lgPos,
			}

			db.index.Tree.ReplaceOrInsert(itm)
		}
	}

}

// 主要是测试用, 删除当前的数据文件重新开始测试
func (db *DB) TrucDB() bool {
	err := os.Remove(filepath.Join(db.option.DirPath, "00000001.data"))
	if err != nil {
		return false
	}
	return true
}

func (db *DB) GetIndex() *Btree {
	return db.index
}
