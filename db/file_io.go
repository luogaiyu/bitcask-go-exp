package db

import (
	"encoding/binary"
	"log"
	"os"
	"path/filepath"
	"unsafe"
)

type FileIO struct {
	DirPath string // 数据路径 exg: 00000001 在定义的数据结构要有注释
}

const (
	tmp_dir = "db/data/"
)

func InitFileIO(Dirpath string) *FileIO {
	return &FileIO{
		DirPath: Dirpath,
	}
}

func (fio *FileIO) Write(lgPos LogPos, logRecord LogRecord) (error, uint32) {
	file, err := os.OpenFile(filepath.Join(fio.DirPath, lgPos.Fid+".data"), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	encLogRecord, n := EncodeLogRecord(logRecord)
	file.WriteAt(encLogRecord, int64(lgPos.Offset))
	defer file.Close()
	return err, n
}

// 需要对 解码方法进行封装
// 20240825: 修复bug: get方法 
func (fio *FileIO) Read(lgPos *LogPos) *LogRecord {
	file, err := os.OpenFile(filepath.Join(fio.DirPath, lgPos.Fid+".data"), os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	cursor := int64(lgPos.Offset)

	lgRecord := &LogRecord{}

	// 查看当前的数据
	rtype := make([]byte, unsafe.Sizeof(lgRecord.RType))
	n, _ := file.ReadAt(rtype, cursor)
	cursor += int64(n)
	lgRecord.RType = rtype[0]

	keySize := make([]byte, unsafe.Sizeof(lgRecord.KeySize))
	n, _ = file.ReadAt(keySize, cursor)
	cursor += int64(n)
	// 20240825 修复bug, 读取出来的 Key Size就是 Key Size
	lgRecord.KeySize = binary.BigEndian.Uint32(keySize)

	valueSize := make([]byte, unsafe.Sizeof(lgRecord.ValueSize))
	n, _ = file.ReadAt(valueSize, cursor)
	cursor += int64(n)
	lgRecord.ValueSize = binary.BigEndian.Uint32(valueSize)

	//获取整体的长度, 为什么需要把keysize,valueSize存在文件中?
	byte_len := uint32(unsafe.Sizeof(lgRecord.RType)) + uint32(unsafe.Sizeof(lgRecord.KeySize)) + uint32(unsafe.Sizeof(lgRecord.ValueSize)) + binary.BigEndian.Uint32(keySize) + binary.BigEndian.Uint32(valueSize)
	value := make([]byte, byte_len)
	n, _ = file.ReadAt(value, int64(lgPos.Offset))
	lgRecord = DecodeLogRecord(value)
	// 返回
	return lgRecord
}
