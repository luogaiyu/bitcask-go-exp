package utils

import (
	"encoding/binary"
)

func UInt32_2_Bytes(n uint32) []byte {
	bs := make([]byte, 4) // int 在 Go 中是 4 个字节（32 位）
	binary.BigEndian.PutUint32(bs, n)
	return bs
}

// 将4字节byte 转换成int64
func Bytes_2_UInt32(value []byte) uint32 {
	res := binary.BigEndian.Uint32(value)
	return res
}
