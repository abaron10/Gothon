package main

import (
	"fmt"
)

type Test struct {
	id string
}

func main() {

	a := []Test{{id: "qqqq"}, {id: "AAAA"}}
	//b := []Test{{id: "1111"}, {id: "kdkdkd"}}
	//c := []int{2, 3, 4, 5}

	//fmt.Println(Pop(&c, -1) + 5)
	//z := Extend(a, b)

	fmt.Println(Pop(&a, -1))
	//c, ok := ExtendInterface(a, b).([]Test)

	//fmt.Println(reflect.TypeOf(Extend(a, b)))

}
