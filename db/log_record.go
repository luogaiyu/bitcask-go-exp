package db

import (
	"encoding/binary"
	// "log"
	"unsafe"
)

// todo : 从上层的看的时候需要知道offset 和size 是否有必要
type LogPos struct {
	Fid    string
	Offset int64
	Size   int64
}
type RType = byte

const (
	LogRecordNormal  RType = iota // 正常记录
	LogRecordDeleted              // 被删除的记录
)

const (
	MaxUintLen32 = 4
)

const maxLogRecordSize = 1 + MaxUintLen32*2

func InitLogPos() *LogPos {
	return &LogPos{
		Fid:    "00000001",
		Offset: 0,
		Size:   0,
	}
}

type LogRecord struct {
	RType     uint8 // 记录类型, 分为0和1
	KeySize   uint32
	ValueSize uint32
	Key       []byte
	Value     []byte
}

// 需要将对应的数据转换成[] byte, 自己来定义编码过程能够少引入很多的无用过程,自己控制逻辑
//
//	编码  rtype :1 + keySize:unit32 + ValueSize:32 + key + value:变长
func EncodeLogRecord(lgRecord LogRecord) ([]byte, uint32) {
	cursor := uint32(0)
	buf := make([]byte, maxLogRecordSize)
	buf[0] = lgRecord.RType
	cursor += uint32(unsafe.Sizeof(lgRecord.RType))

	size := uint32(unsafe.Sizeof(lgRecord.KeySize))
	binary.BigEndian.PutUint32(buf[cursor:cursor+size], lgRecord.KeySize)
	cursor += size

	size = uint32(unsafe.Sizeof(lgRecord.ValueSize))
	binary.BigEndian.PutUint32(buf[cursor:cursor+size], lgRecord.ValueSize)
	cursor += (size + uint32(len(lgRecord.Key)) + uint32(len(lgRecord.Value)))

	encBytes := make([]byte, cursor)
	copy(encBytes[:maxLogRecordSize], buf[:maxLogRecordSize])
	copy(encBytes[maxLogRecordSize:maxLogRecordSize+uint32(len(lgRecord.Key))], lgRecord.Key)
	copy(encBytes[maxLogRecordSize+uint32(len(lgRecord.Key)):maxLogRecordSize+uint32(len(lgRecord.Key))+uint32(len(lgRecord.Value))], lgRecord.Value)
	return encBytes, cursor

}

// 将对应的[]byte 转换成 logRecord
func DecodeLogRecord(value []byte) LogRecord {
	lgRecord := LogRecord{}
	var index = 0
	rType, size := binary.Varint(value[index:])
	index += size
	lgRecord.RType = byte(rType)
	lgRecord.Value = value[index:]
	return lgRecord
}
