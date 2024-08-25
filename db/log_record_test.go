package db

import (
	"fmt"
	"os"
	"testing"
)

func TestEncodeLogRecord(t *testing.T) {
	file, err := os.OpenFile("/Users/bowen.yin/code/kv-projects/bitcask-go-exp/db/data/00000001.data", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	lgRecord := &LogRecord{
		RType:     0, // 记录类型, 分为0和1
		KeySize:   uint32(200),
		ValueSize: uint32(200),
		Key:       []byte("12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
		Value:     []byte("12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"),
	}
	encLogRecord, _ := EncodeLogRecord(*lgRecord)

	n, err := file.WriteAt(encLogRecord, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	defer file.Close()
}

func TestDecodeLogRecord(t *testing.T){
	file, _ := os.OpenFile("/Users/bowen.yin/code/kv-projects/bitcask-go-exp/db/data/00000001.data", os.O_RDONLY, 0666)
	b := make([]byte, 409)
	_, _ = file.ReadAt(b, 0)
	resLogRecord := DecodeLogRecord(b)
	fmt.Println(resLogRecord)
	defer file.Close()
}
