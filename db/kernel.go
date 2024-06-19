// 首先需要将代码的流程跑通
package db

import (
	"bytes"
	"encoding/gob"
	"io"
	"log"
	"os"
	// "fmt"
)

type Log struct {
	Key   []byte
	Value []byte
}

// bw DB v1
func Put(key []byte, value []byte) {
	file, err := os.OpenFile(kernel_data_path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	lg := Log{
		Key:   key,
		Value: value,
	}
	enc := gob.NewEncoder(file)
	err = enc.Encode(&lg)
	if err != nil {
		log.Fatal(err)
	}
}

func Get(key []byte) (Log,error) {
	file, err := os.Open(kernel_data_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lg Log
	for {
		dec := gob.NewDecoder(file)
		err = dec.Decode(&lg)
		if err == io.EOF {
			return lg,err
		}
		if(bytes.Equal(lg.Key,key)){
			return lg,nil
		}
	}
	
}

// 使用文件来保存对应的记录
func Init() {
	os.MkdirAll(kernel_dir_path, 0755)
	// 打开文件，如果文件不存在则创建
	file, err := os.OpenFile(kernel_data_path, os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

// 测试用,清空当前的数据目录
func Truc() {
	err := os.RemoveAll(kernel_dir_path)
	if err != nil {
		log.Fatal(err)
	}
}
