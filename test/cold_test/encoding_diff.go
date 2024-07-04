package test

import (
	"bitcask-go-exp/db"
	"encoding/gob"
	"fmt"
	"os"
	"testing"
)

// 用来衡量 为什么需要自定义编码过程

// // FileIO 标准系统文件 IO
// type FileIO struct {
// 	fd *os.File // 系统文件描述符
// }

// EncodeLogRecord 对 LogRecord 进行编码，返回字节数组及长度
// type: 1字节
// value: 1字节

// crc type keySize valueSize
// 4 +  1  +  5   +   5 = 15

func A() {
	lgReocrd := db.LogRecord{
		Value: []byte("1"),
		RType: db.LogRecordNormal,
	}

	file, err := os.OpenFile(
		"/Users/bowen.yin/code/kv-projects/bitcask-go-exp/test/data/test.data",
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0644,
	)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}
	enc := gob.NewEncoder(file)
	err = enc.Encode(lgReocrd)

	if err != nil {
		fmt.Println(err)
	}
}

func B() {
	lgReocrd := db.LogRecord{
		Value: []byte("1"),
		RType: db.LogRecordNormal,
	}
	file, err := os.OpenFile(
		"/Users/bowen.yin/code/kv-projects/bitcask-go-exp/test/data/test_B.data",
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0644,
	)
	defer file.Close()

	encBytes, _ := db.EncodeLogRecord(lgReocrd)

	fmt.Println(len(encBytes))
	_, err = file.Write(encBytes)
	if err != nil {
		fmt.Println(err)
	}

}

func TestA(t *testing.T) {
	A()
}

func TestB(t *testing.T) {
	B()
}
