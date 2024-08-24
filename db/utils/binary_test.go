package utils

import (
	"testing"
	"fmt"
)

func TestInt642Bytes(t *testing.T) {
	fmt.Println(Bytes_2_UInt32(UInt32_2_Bytes(10))) 
}
