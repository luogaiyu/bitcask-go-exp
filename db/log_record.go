package db

import (
	"encoding/binary"
	// "log"
	"unsafe"
)

// todo : 从上层的看的时候需要知道size 是否有必要
type LogPos struct {
	Fid    string
	Offset uint64
	// Size   int64
}

const (
	MaxUintLen32 = 4
)

const maxLogRecordSize = 1 + MaxUintLen32*2

func InitLogPos() *LogPos {
	return &LogPos{
		Fid:    "00000001",
		Offset: 0,
	}
}

type RType = uint8

const (
	LogRecordNormal  RType = iota // 正常记录
	LogRecordDeleted              // 被删除的记录
)

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
//
// 方法需要加上 err异常判断
// 20240825 修复bug, unsafe.SizeOf 
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

// 将对应的[]byte 转换成 logRecord 这里需要把bool也返回
func DecodeLogRecord(value []byte) LogRecord {
	lgRecord := LogRecord{}
	var index = 0
	RType := uint8(value[index])
	index += 1
	lgRecord.RType = RType

	KeySize := binary.BigEndian.Uint32(value[index : index+4]) // 这个地方后面可以暴露出去, 使用配置项来进行控制
	index += 4
	lgRecord.KeySize = KeySize

	ValueSize := binary.BigEndian.Uint32(value[index : index+4])
	index += 4
	lgRecord.ValueSize = ValueSize

	Key := value[index : index+int(KeySize)]
	index += int(KeySize)
	lgRecord.Key = Key

	Value := value[index : index+int(ValueSize)]
	index += int(ValueSize)
	lgRecord.Value = Value

	return lgRecord
}
