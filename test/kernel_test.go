package test

import (
	"bitcask-go-exp/db"
	"fmt"
	"log"
	"testing"
)

func TestInit(t *testing.T) {
    db.Init()
}

func TestPut(t *testing.T) {
	k := []byte("key")
	v := []byte("value")
    db.Put(k,v)
	fmt.Println()
}

func TestGet(t *testing.T) {
    db.Init()
	k := []byte("key")
	lg, err:= db.Get(k)
	if(err != nil) {
		log.Fatal("没有数据")
	}else{
		fmt.Println(string(lg.Key))
		fmt.Println(string(lg.Value))
	}
}


func TestTruc(t *testing.T) {
	db.Truc()
}