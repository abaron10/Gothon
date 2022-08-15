package main

import "fmt"

func main() {

	//a := []int{1, 2, 4}
	b := []string{"aaa", "bbb", "ccc"}
	c := PopInt64(&[]int64{1, 2, 3}, 1)

	fmt.Println(c)
	fmt.Println(b)

}
