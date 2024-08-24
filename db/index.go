package db

import (
	"bytes"
	"errors"

	"github.com/google/btree"
)

// btree 能存 能取 能删
type Btree struct {
	Tree *btree.BTree
}

type Item struct {
	key   []byte
	lgPos *LogPos
}

func GetNewItemByKey(key []byte) *Item {
	return &Item{
		key: key,
	}

}
func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}

func (ai *Item) GetKey(bi btree.Item) []byte {
	return ai.key
}
// 需要返回对应的引用
func  InitBTree() *Btree {
	return &Btree{
		Tree: btree.New(32),
	}
}

func (bt *Btree) Put(key []byte, lgPos LogPos) {
	itm := &Item{
		key:   key,
		lgPos: &lgPos,
	}
	bt.Tree.ReplaceOrInsert(itm)
}

func (bt *Btree) Get(key []byte) (btree.Item,error) {
	itm := &Item{
		key: key,
	}
	res := bt.Tree.Get(itm)
	if res != nil {
		return bt.Tree.Get(itm),nil
	}else {
		return nil, errors.New("the key in index not found, value can't get ")
	}
	
}

func (bt *Btree) Delete(key []byte) {
	itm := &Item{
		key: key,
	}
	bt.Tree.Delete(itm)
}
