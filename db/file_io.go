package db

import (
	"path/filepath"
	// "encoding/binary"
	"log"
	"os"
	"unsafe"
)

type FileIO struct {
	DirPath string // 数据路径
}
const (
	tmp_dir = "db/data/"
)

func InitFileIO(Dirpath string) *FileIO{
	return &FileIO{
		DirPath: Dirpath,
	}
}

func (fio *FileIO) Write(lgPos LogPos, logRecord LogRecord) (error,uint32) {
	file, err := os.OpenFile(filepath.Join(fio.DirPath,lgPos.Fid+".data") , os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	encLogRecord, n := EncodeLogRecord(logRecord)
	file.WriteAt(encLogRecord, int64(lgPos.Offset))
	defer file.Close()
	return err,n
}


func (fio *FileIO) Read(lgPos *LogPos) *LogRecord {
	file, err := os.OpenFile(filepath.Join(fio.DirPath,lgPos.Fid+".data"), os.O_RDONLY, 0666)
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
	lgRecord.KeySize = uint32(len(keySize))


	valueSize := make([]byte, unsafe.Sizeof(lgRecord.ValueSize))
	n, _ = file.ReadAt(valueSize, cursor)
	cursor += int64(n)
	lgRecord.ValueSize = uint32(len(keySize))

	key := make([]byte, lgRecord.KeySize)
	n, _ = file.ReadAt(key, cursor)
	cursor += int64(n)
	lgRecord.Key = key

	value := make([]byte, lgRecord.ValueSize)
	n, _ = file.ReadAt(value, cursor)
	cursor += int64(n)
	lgRecord.Value = value

	return lgRecord
}
