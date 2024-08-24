package db

import (
	"fmt"
	"testing"

	"github.com/google/btree"
)
func TestPut(t *testing.T){
	// fid string, offset int64, size int64
	key := []byte("1234")
	fid := "./data/" + "00000001" +".data"
	lgPos := &LogPos{
		Fid:fid,
		Offset:1,
		// Size:1,
	}
	bt := btree.New(2)
	btree := Btree{
		Tree:bt,
	}
	btree.Put(key,lgPos)
	// Put(key,)
}



func TestGet(t *testing.T){
	// fid string, offset int64, size int64
	key := []byte("1234")
	fid := "./data/" + "00000001" +".data"
	lgPos := &LogPos{
		Fid:fid,
		Offset:1,
		// Size:1,
	}
	bt := btree.New(2)
	btree := Btree{
		Tree:bt,
	}
	btree.Put(key,lgPos)
	res, err := btree.Get(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res.(*Item).lgPos)
	// Put(key,)
}
