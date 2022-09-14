package gothonSlice

//the Extend() method adds the specified list elements to the end of the current list. Requires two slices and returns a new
//slice with the extesion of both
func Extend[T any](a []T, b []T) []T {
	for _, v := range b {
		a = append(a, v)
	}
	return a
}

//the Pop() method removes the element at the specified position, It requires a pointer to the
//slice and the index to be removed. Changes are made in place with the given slice and the method returns the popped value.
//Supports negative indexing e.i -1 corresponds to the last position of a slice
func Pop[T any](a *[]T, index int) T {
	if emptySlice(len((*a))) {
		panic("IndexError: pop from empty list")
	}
	index = validateIndex(index, len(*a))
	popped := (*a)[index]
	(*a) = append((*a)[:index], (*a)[index+1:]...)
	return popped
}

//the Index() method returns the position at the first occurrence of the specified value.It requires an slice and a value
//to be checked inside the slice. Returns an integer with the index of the element, if the value doesn't exist return -1
func Index[T comparable](a []T, value T) int {
	for i, v := range a {
		if v == value {
			return i
		}
	}
	return -1
}

//the Contains() method check if the item exists in the slice. It requires an slice and a value to be checked inside
//the slice. Returns a boolean whether the element exists or not
func Contains[T comparable](a []T, value T) bool {
	return Index(a, value) != -1
}

//the Reverse() method reverses the sorting order of the elements. Receives a pointer from the given slice
//Changes are made in place with the given slice.
func Reverse[T any](a *[]T) {
	r := []T{}
	for i := len((*a)) - 1; i >= 0; i-- {
		r = append(r, (*a)[i])
	}
	(*a) = r
}

//the Remove() method removes the first occurrence of the element with the specified value. It requires a pointer to the
//slice and the raw value to be removed from the slice. Changes are made in place with the given slice
func Remove[T comparable](a *[]T, value T) {
	if index := Index((*a), value); index != -1 {
		Pop(a, Index((*a), value))
	}
}

//the Insert() method inserts the specified value at the specified position. It requires a pointer to the slice, an index,
// and a value to be inserted. Changes are made in place with the given slice
func Insert[T any](a *[]T, index int, value T) {
	if len((*a)) <= index { // nil or empty slice or after last element
		(*a) = append((*a), value)
		return
	}
	index = validateIndex(index, len(*a))
	(*a) = append((*a)[:index+1], (*a)[index:]...) // index < len(a)
	(*a)[index] = value
}

//the Clear() method removes all the elements from the slice
func Clear[T any](a *[]T) {
	(*a) = []T{}
}

//the Count() method returns the number of elements with the specified passed value
func Count[T comparable](a []T, value T) int {
	c := 0
	for _, v := range a {
		if v == value {
			c++
		}
	}
	return c
}

//the Copy() method returns a copy of the specified slice
func Copy[T any](a []T) []T {
	var copy []T
	copy = append(copy, a...)
	return copy
}

//helpers methods

func binarySearch(a []float64, x float64) int {
	r := -1 // not found
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == x {
			r = mid // found
			break
		} else if a[mid] < x {
			start = mid + 1
		} else if a[mid] > x {
			end = mid - 1
		}
	}
	return r
}

func validateIndex(index int, size int) int {
	if index > size-1 {
		panic("IndexError: pop index out of range")
	} else if index < 0 {
		index = negativeIndexTranslation(size, index)
	}
	return index
}

func negativeIndexTranslation(listSize int, index int) int {
	t := listSize + index
	if t < 0 {
		panic("IndexError: Index out of range")
	}
	return t
}

func emptySlice(a int) bool {
	if a == 0 {
		return true
	}
	return false
}
