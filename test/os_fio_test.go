package test

import (
	"fmt"
	"testing"
)

// todo : 从上层的看的时候需要知道size 是否有必要
type tmp_object struct {
	Fid    string
}



func C() {
	res_map := make(map[string]*tmp_object)
	to := &tmp_object{
		Fid: "1",
	}
	res_map["1"] = to
	fmt.Println(res_map["1"])
	to.Fid = "2"
	fmt.Println(res_map["1"])

	
}

func TestC(t *testing.T) {
	C()
}
