package main

import "fmt"

type Test struct {
	id string
}

func main() {
	a := []int64{1, 2, 3}
	Insert(&a, -2, 9)
	fmt.Println(a)
}
