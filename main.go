package main

import (
	"fmt"
	"github.com/abaron10/Gothon/gothonHeapq"
)

func main() {

	a := []int{9, 2, 3, 7, 100, 64, 1004, 2}

	gothonHeapq.Heapify(&a)

	gothonHeapq.HeapPush(&a, 1)

	fmt.Println(a)
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
