package main

import (
	"fmt"
	"github.com/abaron10/Gothon/gothonSlice"
)

func main() {

	a := []int32{1, 2, 3}
	b := gothonSlice.Copy(a)
	gothonSlice.Pop(&b, len(b)-1)
	fmt.Println(a, &b)

}

func Test(a interface{}) {
	anySlice, ok := a.([]interface{})
	fmt.Println(anySlice, ok)
	switch x := a.(type) {
	case []string:
		fmt.Println("slcie", x)
	case string:
		fmt.Println("string")

	default:
		fmt.Println("None")
	}

}
