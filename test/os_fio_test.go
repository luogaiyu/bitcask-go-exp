package test

import (
	"testing"
	"fmt"
	"unsafe"
)

// todo : 从上层的看的时候需要知道size 是否有必要
type tmp_object struct {
	Fid    string
}


type LogRecord_tmp struct {
	RType     uint8 // 记录类型, 分为0和1
	KeySize   uint32
	ValueSize uint32
	Key       []byte
	Value     []byte
}


func C() {
	test_str := []byte("1234")
	// fmt.Println(len(test_str))
	// res:= len(test_str)
	fmt.Println(unsafe.Sizeof(test_str))
	// fmt.Println(res)
}

func TestC(t *testing.T) {
	C()
}
