package collection

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	Name string
	Age  int
}

func TestInterfaceToString(t *testing.T) {
	fmt.Println(interfaceToString(1))
}

func (ts TestStruct) String() string {
	return "here"
}
