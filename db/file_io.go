package db

import (
	"encoding/binary"
	"log"
	"os"
	"unsafe"
)

type FileIO struct {
}
const (
	tmp_dir = "./data/"
)

func InitFileIO() *FileIO{
	return &FileIO{}
}

func (fio *FileIO) Write(lgPos LogPos, logRecord LogRecord) error {
	file, err := os.OpenFile(tmp_dir + lgPos.Fid+".data", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	encLogRecord,_ := EncodeLogRecord(logRecord)
	file.WriteAt(encLogRecord, lgPos.Offset)
	defer file.Close()
	return err
}


func (fio *FileIO) Read(lgPos *LogPos) *LogRecord {
	file, err := os.OpenFile(tmp_dir + lgPos.Fid+".data", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	cursor := lgPos.Offset
	lgRecord := &LogRecord{}

	// 查看当前的数据
	rtype := make([]byte, unsafe.Sizeof(lgRecord.RType))	
	n, _ := file.ReadAt(rtype, cursor)
	cursor += int64(n)
	lgRecord.RType = rtype[0]
	
	keySize := make([]byte, unsafe.Sizeof(lgRecord.KeySize))
	n, _ = file.ReadAt(keySize, cursor)
	cursor += int64(n)
	lgRecord.KeySize = binary.BigEndian.Uint32(keySize)


	valueSize := make([]byte, unsafe.Sizeof(lgRecord.ValueSize))
	n, _ = file.ReadAt(valueSize, cursor)
	cursor += int64(n)
	lgRecord.ValueSize = binary.BigEndian.Uint32(valueSize)

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
