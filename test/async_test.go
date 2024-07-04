package test


import 
(
	"testing"
	"fmt"
)

type m struct{
	k string
	v string
}

func A(){
	t :=  m{
		k: "123",
		v: "321",
	} 
	t1 := &t 
	t1.k = "1234"

	fmt.Println(t1)
	fmt.Println(t)
}


func TestA(t *testing.T){
	A()
}