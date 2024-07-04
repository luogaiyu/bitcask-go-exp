package db

import (
	"encoding/binary"
	"fmt"
	"os"
	"unsafe"
)

// 注意体会go语言的 值传递和引用传递
type DB struct {
	index   *Btree // 保存当前数据库的索引
	logFile *LogPos //保存当前数据库对应文件日志的地址, db不需要维护这个pos, 这个pos是文件系统需要的
	fio     *FileIO // 用来存放一些和文件交互的逻辑
}


// 初始化的方法
// merge的时候 才会新建索引的hint文件
// DataFile
func InitDB() *DB {
	// loadIndexFromFile() 从文件中重建索引
	return &DB{
		index: InitBTree(),
		logFile: InitLogPos(),
		fio: InitFileIO(),
	}

}

func (db *DB) Put(key []byte, value []byte) {
	// 首先将数据保存到索引中, 再将数据写入到内存中, 为什么不将所有的数据写入到内存中, 如果全部写入到内存中可以存多少条数据?
	db.index.Put(key, db.logFile)
	// db.fio.Write(db.logFile,value)
	db.logFile.Size = int64(len(value))
	db.logFile.Offset = db.logFile.Offset + db.logFile.Size
}

func (db *DB) Get(key []byte) []byte {
	itm := db.index.Get(key)
	return db.fio.Read(itm.(*Item).lgPos).Value

}
func (db *DB) Delete(key []byte, value []byte) {
	db.index.Delete(key)
	// logRecord = &LogRecord{
	// 	value : value,
	// 	rType: LogRecordNormal,
	// }
	// db.fio.Write()
}

// 从数据文件中构建索引
func (db *DB) loadIndexFromDataFile() {
	file, _ := os.OpenFile("./data/00000001.data", os.O_RDONLY, 0666)

	lgRecord := LogRecord{}
	lgPos := &LogPos{}
	cursor := int64(0)
	
	for {
		lgPos= &LogPos{
			Fid: "00000001",
		}
		// 一个个 record读取
		lgPos.Offset = cursor
		rtype := make([]byte, unsafe.Sizeof(lgRecord.RType))
		
		n, _ := file.ReadAt(rtype, cursor)
		if n == 0 {
			break
		}
		cursor += int64(n)
		
		keySize := make([]byte, unsafe.Sizeof(lgRecord.KeySize))

		n, _ = file.ReadAt(keySize, cursor)
		cursor += int64(n)
		
		valueSize := make([]byte, unsafe.Sizeof(lgRecord.ValueSize))
		n, _ = file.ReadAt(valueSize, cursor)
		cursor += int64(n)

		key := make([]byte, binary.BigEndian.Uint32(keySize))
		n,_ = file.ReadAt(key, cursor)
		cursor += int64(n)


		value := make([]byte, binary.BigEndian.Uint32(valueSize))
		n,_ = file.ReadAt(value, cursor)
		cursor += int64(n)

		lgPos.Size = (cursor - lgPos.Offset)
		// 设置新的pos
		itm := &Item{
			key : key,
			lgPos : lgPos,
		}
		
		db.index.Tree.ReplaceOrInsert(itm)
	}
	
	// 再按照字节 把文件依次读出来
	// 放到新的树中
	
}

func (db *DB) syntax() {
	fmt.Println(db.index)
}

