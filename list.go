package main

import (
	"fmt"
	"reflect"
)

func interfaceSliceToIntSlice(r []int) []int {
	return r

}

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------HELPERS METHOD------------------------------------------------
---------------------------------------------------------------------------------------------------------
*/

func negativeIndexTranslation(listSize int, index int) int {
	t := listSize + index
	if t < 0 {
		panic("IndexError: pop index out of range")
	}
	return t
}

func emptySlice(a int) bool {
	if a == 0 {
		return true
	}
	return false
}

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------EXTENDS METHOD--------------------------------------------------
-----------------------------------------------------------------------------------------------------------
The extend() method adds the specified list elements (or any iterable) to the end of the current list.-----
Ref https://www.w3schools.com/python/ref_list_extend.asp --------------------------------------------------
*/

func ExtendInt(a []int, b []int) []int {
	for _, v := range b {
		a = append(a, v)
	}
	return a
}

func ExtendInt32(a []int32, b []int32) []int32 {
	for _, v := range b {
		a = append(a, v)
	}
	return a
}

func ExtendInt64(a []int64, b []int64) []int64 {
	for _, v := range b {
		a = append(a, v)
	}
	return a
}

func ExtendFloat32(a []float32, b []float32) []float32 {
	for _, v := range b {
		a = append(a, v)
	}
	return a
}

func ExtendFloat64(a []float64, b []float64) []float64 {
	for _, v := range b {
		a = append(a, v)
	}
	return a
}

func ExtendString(a []string, b []string) []string {
	for _, v := range b {
		a = append(a, v)
	}
	return a
}

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------POP METHOD------------------------------------------------------
-----------------------------------------------------------------------------------------------------------
The pop() method removes the element at the specified position.--------------------------------------------
Ref https://www.w3schools.com/python/ref_list_pop.asp------------------------------------------------------
*/

func PopInt(a []int, index int) ([]int, int) {
	if emptySlice(len(a)) {
		panic("IndexError: pop from empty list")
	}
	r := []int{}
	if index > len(a)-1 {
		panic("IndexError: pop index out of range")
	} else if index < 0 {
		index = negativeIndexTranslation(len(a), index)
	}
	for i, v := range a {
		if i != index {
			r = append(r, v)
		}
	}
	return r, a[index]
}

func PopInt32(a []int32, index int) ([]int32, int32) {
	if emptySlice(len(a)) {
		panic("IndexError: pop from empty list")
	}
	r := []int32{}
	if index > len(a)-1 {
		panic("IndexError: pop index out of range")
	} else if index < 0 {
		index = negativeIndexTranslation(len(a), index)
	}
	for i, v := range a {
		if i != index {
			r = append(r, v)
		}
	}
	return r, a[index]
}

func PopInt64(a []int64, index int) ([]int64, int64) {
	if emptySlice(len(a)) {
		panic("IndexError: pop from empty list")
	}
	r := []int64{}
	if index > len(a)-1 {
		panic("IndexError: pop index out of range")
	} else if index < 0 {
		index = negativeIndexTranslation(len(a), index)
	}
	for i, v := range a {
		if i != index {
			r = append(r, v)
		}
	}
	return r, a[index]
}

func PopFloat32(a []float32, index int) ([]float32, float32) {
	if emptySlice(len(a)) {
		panic("IndexError: pop from empty list")
	}
	r := []float32{}
	if index > len(a)-1 {
		panic("IndexError: pop index out of range")
	} else if index < 0 {
		index = negativeIndexTranslation(len(a), index)
	}
	for i, v := range a {
		if i != index {
			r = append(r, v)
		}
	}
	return r, a[index]
}

func PopFloat64(a []float64, index int) ([]float64, float64) {
	if emptySlice(len(a)) {
		panic("IndexError: pop from empty list")
	}
	r := []float64{}
	if index > len(a)-1 {
		panic("IndexError: pop index out of range")
	} else if index < 0 {
		index = negativeIndexTranslation(len(a), index)
	}
	for i, v := range a {
		if i != index {
			r = append(r, v)
		}
	}
	return r, a[index]
}

func PopString(a []string, index int) ([]string, string) {
	if emptySlice(len(a)) {
		panic("IndexError: pop from empty list")
	}
	r := []string{}
	if index > len(a)-1 {
		panic("IndexError: pop index out of range")
	} else if index < 0 {
		index = negativeIndexTranslation(len(a), index)
	}
	for i, v := range a {
		if i != index {
			r = append(r, v)
		}
	}
	return r, a[index]
}

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------INDEX METHOD----------------------------------------------------
-----------------------------------------------------------------------------------------------------------
The index() method returns the position at the first occurrence of the specified value.--------------------
Ref https://www.w3schools.com/python/ref_list_index.asp----------------------------------------------------
*/

func Index(lA []interface{}, value interface{}) int {
	rType := reflect.TypeOf(value)
	rKind := rType.Kind()
	var response = -1

	if rKind.String() == "string" {
		s := make([]string, len(lA))
		for i, v := range lA {
			s[i] = fmt.Sprint(v)
		}
		response = Search(len(lA), func(i int) bool { return s[i] >= value.(string) })
	} else if rKind.String() == "float64" {
		s := make([]float64, len(lA))
		for i, v := range lA {
			s[i] = v.(float64)
		}
		response = Search(len(lA), func(i int) bool { return s[i] >= value.(float64) })
	} else if rKind.String() == "int" {
		s := make([]string, len(lA))
		for i, v := range lA {
			s[i] = fmt.Sprint(v)
		}
		response = Search(len(lA), func(i int) bool { return s[i] >= value.(string) })

	}
	return response
}

func interfaceToArr(lA []interface{}, s []interface{}) []interface{} {
	for i, v := range lA {
		s[i] = fmt.Sprint(v)
	}
	return s
}

func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}
