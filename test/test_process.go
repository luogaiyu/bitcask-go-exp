// 首先需要将代码的流程跑通
package test

import (
    "os"
    "log"
)
// bw DB v1
func put(data []byte ) {

}

func get(data []byte) {
	
}

func 
func main() {
    // 打开文件，如果文件不存在则创建
    file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // 写入字符串
    _, err = file.WriteString("Hello, World!")
    if err != nil {
        log.Fatal(err)
    }
}