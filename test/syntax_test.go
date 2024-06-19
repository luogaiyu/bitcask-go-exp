package test

import (
	"encoding/gob"
	"fmt"
	// "fmt"
	"log"
	"os"
	"testing"
)


type Log struct {
	key   []byte
	value []byte
}

func A() {
	// 创建一个文件
	file, err := os.Open("/Users/bowen.yin/code/kv-projects/bitcask-go-exp/test/data/kennel.data")
	if err != nil{
		
	}
	fmt.Print(file)
	var lg Log
	dec := gob.NewDecoder(file)
	err = dec.Decode(&lg)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Print(lg)
}

func TestA(t *testing.T) {
	A()
}
