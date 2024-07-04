package db

import
(
	"github.com/google/btree"
	"testing"
)
func TestPut(t *testing.T){
	// fid string, offset int64, size int64
	key := []byte("1234")
	fid := "./data/" + "00000001" +".data"
	lgPos := &LogPos{
		Fid:fid,
		Offset:1,
		Size:1,
	}
	bt := btree.New(2)
	btree := Btree{
		Tree:bt,
	}
	btree.Put(key,lgPos)
	// Put(key,)
}
