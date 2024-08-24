package utils

import
(
	"testing"
	"fmt"
	"time"
)

func TestGetRandomString(T *testing.T){
	fmt.Println(time.Now().UnixNano())
	for i:=0; i < 10; i ++ {
		fmt.Println(GetRandomString())
	}
}