package main

//The Extend() method adds the specified list elements to the end of the current list
func Extend[T any](a []T, b []T) []T {
	for _, v := range b {
		a = append(a, v)
	}
	return a
}

//The Pop() method removes the element at the specified position
func Pop[T any](a *[]T, index int) T {
	if emptySlice(len((*a))) {
		panic("IndexError: pop from empty list")
	}
	index = validateIndex(index, len(*a))
	popped := (*a)[index]
	(*a) = append((*a)[:index], (*a)[index+1:]...)
	return popped
}

//The Index() method returns the position at the first occurrence of the specified value
func Index[T comparable](a []T, value T) int {
	for i, v := range a {
		if v == value {
			return i
		}
	}
	return -1
}

//The Contains() method check if the item exists in the list
func Contains[T comparable](a []T, value T) bool {
	return Index(a, value) != -1
}

//The Reverse() method reverses the sorting order of the elements. Receives a pointer from the given slice
func Reverse[T any](a *[]T) {
	r := []T{}
	for i := len((*a)) - 1; i >= 0; i-- {
		r = append(r, (*a)[i])
	}
	(*a) = r
}

//The Remove() method removes the first occurrence of the element with the specified value.
func Remove[T comparable](a *[]T, value T) {
	Pop(a, Index((*a), value))
}

//The Insert() method inserts the specified value at the specified position.
func Insert[T any](a *[]T, index int, value T) {
	index = validateIndex(index, len(*a))
	if len((*a)) == index { // nil or empty slice or after last element
		(*a) = append((*a), value)
	}
	(*a) = append((*a)[:index+1], (*a)[index:]...) // index < len(a)
	(*a)[index] = value
}

//Helpers methods
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
