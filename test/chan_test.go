package test

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)


// Chan
func TestChan(T *testing.T)  {
	// 创建一个缓冲区大小为 2 的通道
	// errChan := make(chan error, 2)

	assert := assert.New(T)
	test := assert.Equal(1,1)

	fmt.Println(test)
	
	test = assert.Equal(1,2)
	fmt.Println(test)
		
}