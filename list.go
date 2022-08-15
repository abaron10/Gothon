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

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------EXTENDS METHOD------------------------------------------------
---------------------------------------------------------------------------------------------------------
*/
func ExtendInt(a interface{}, b interface{}) []int {
	if s1, ok := a.([]int); ok {
		if s2, ok := b.([]int); ok {
			for _, v := range s2 {
				s1 = append(s1, v)
			}
			return s1
		}
		panic(fmt.Sprintf("TypeError: can only concatenate []int32 not %s", reflect.TypeOf(b)))
	}
	panic(fmt.Sprintf("TypeError: can only concatenate []int32 not %s", reflect.TypeOf(a)))
}

func ExtendInt32(a interface{}, b interface{}) []int32 {
	if s1, ok := a.([]int32); ok {
		if s2, ok := b.([]int32); ok {
			for _, v := range s2 {
				s1 = append(s1, v)
			}
			return s1
		}
		panic(fmt.Sprintf("TypeError: can only concatenate []int32 not %s", reflect.TypeOf(b)))
	}
	panic(fmt.Sprintf("TypeError: can only concatenate []int32 not %s", reflect.TypeOf(a)))
}

func ExtendInt64(a interface{}, b interface{}) []int64 {
	if s1, ok := a.([]int64); ok {
		if s2, ok := b.([]int64); ok {
			for _, v := range s2 {
				s1 = append(s1, v)
			}
			return s1
		}
		panic(fmt.Sprintf("TypeError: can only concatenate []int64 not %s", reflect.TypeOf(b)))
	}
	panic(fmt.Sprintf("TypeError: can only concatenate []int64 not %s", reflect.TypeOf(a)))
}

func ExtendFloat32(a interface{}, b interface{}) []float32 {
	if s1, ok := a.([]float32); ok {
		if s2, ok := b.([]float32); ok {
			for _, v := range s2 {
				s1 = append(s1, v)
			}
			return s1
		}
		panic(fmt.Sprintf("TypeError: can only concatenate []float32 not %s", reflect.TypeOf(b)))
	}
	panic(fmt.Sprintf("TypeError: can only concatenate []float32 not %s", reflect.TypeOf(a)))
}

func ExtendString(a interface{}, b interface{}) []string {
	if s1, ok := a.([]string); ok {
		if s2, ok := b.([]string); ok {
			for _, v := range s2 {
				s1 = append(s1, v)
			}
			return s1
		}
		panic(fmt.Sprintf("TypeError: can only concatenate []float32 not %s", reflect.TypeOf(b)))
	}
	panic(fmt.Sprintf("TypeError: can only concatenate []float32 not %s", reflect.TypeOf(a)))
}

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------POP METHOD------------------------------------------------
---------------------------------------------------------------------------------------------------------
*/
func PopInt(a []int, index int) ([]int, int) {
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

func PopInt32(a []int, index int) ([]int, int) {
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
