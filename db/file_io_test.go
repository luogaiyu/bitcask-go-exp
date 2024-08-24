package db

import (
	// "bytes"
	// "fmt"
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	fid := "00000001"
	offset := int64(0)

	key := []byte("1234")
	value := []byte("4321")
	lgReocrd := LogRecord{
		RType:     LogRecordNormal,
		KeySize:   uint32(len(key)),
		ValueSize: uint32(len(value)),
		Key:       key,
		Value:     value,
	}
	lgPos := LogPos{
		Fid:    fid,
		Offset: uint64(offset),
		// Size: int64(1+4+4+len(lgReocrd.Key)+len(lgReocrd.Value)),
	}
	fileIO := FileIO{}
	fileIO.Write(lgPos, lgReocrd)
}

func TestRead(t *testing.T) {
	fid := "00000001"
	offset := int64(1)
	// size := int64(9)
	lgPos := &LogPos{
		Fid:    fid,
		Offset: uint64(offset),
		// Size: size,
	}
	fileIO := FileIO{}
	res := fileIO.Read(lgPos)

	fmt.Println(string(res.Key))

	fmt.Println(string(res.Value))
}
