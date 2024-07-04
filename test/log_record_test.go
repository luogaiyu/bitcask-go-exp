package test

import (
	"bitcask-go-exp/db"
	"fmt"
	"testing"
)



func B() {
	lgReocrd := db.LogRecord{
		Value: []byte("12345"),
		RType: db.LogRecordNormal,
	}
	encBytes, _ := db.EncodeLogRecord(lgReocrd)
	decLgRecord := db.DecodeLogRecord(encBytes)
	fmt.Println(decLgRecord.RType)
	fmt.Println(string(decLgRecord.Value))
}

func TestB(t *testing.T) {
	B()
}
