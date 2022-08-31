package gothonHeapq

import (
	"github.com/abaron10/Gothon/gothonSlice"
	"golang.org/x/exp/constraints"
)

//This interface restrict variables that can be set in generics.
type Number interface {
	constraints.Integer | constraints.Float
}

type Heap[T Number] struct {
	arr  []T
	size int
}

//Transform list x into a heap, in-place.
//Receives a pointer from a slice.Changes are made in place
func Heapify[T Number](arr *[]T) {
	heap := Heap[T]{arr: []T{}}
	for _, val := range *arr {
		heap.insert(val)
	}
	*arr = heap.arr
}

//Push the value item onto the heap, maintaining the heap invariant.
//Receives a pointer from a slice and a value to be pushed.
//Changes are made in place
func HeapPush[T Number](arr *[]T, val T) {
	*arr = append(*arr, val)
	Heapify(arr)
}

//Push item on the heap, then pop and return the smallest item from the heap.
//Receives a pointer from a slice and a value to be pushed.
//Changes are made in place
func HeapPushPop[T Number](arr *[]T, val T) T {
	HeapPush(arr, val)
	return HeapPop(arr)
}

//Pop and return the smallest item from the heap, maintaining the heap invariant.
//Receives a pointer from a slice
//Changes are made in place
func HeapPop[T Number](arr *[]T) T {
	lastItem := gothonSlice.Pop(arr, -1)
	if len(*arr) == 0 {
		return lastItem
	}
	smallerItem := (*arr)[0]
	gothonSlice.Remove(arr, smallerItem)
	gothonSlice.Insert(arr, 0, lastItem)
	heap := Heap[T]{arr: (*arr)}
	heap.bubbleDown()
	*arr = heap.arr
	return smallerItem
}

//helpers methods
func (h *Heap[T]) insert(value T) {
	h.arr = append(h.arr, value)
	h.size++
	h.bubbleUp()
}

func (h *Heap[T]) bubbleUp() {
	index := h.size - 1
	for index > 0 && h.arr[index] < h.arr[parent(index)] {
		h.swap(index, parent(index))
		index = parent(index)
	}
}

func (h *Heap[T]) bubbleDown() {
	parent := 0
	left, right := calculateChildren(parent)
	for validBubbleDown(left, right, parent, h.arr) {
		smaller := getSmallerChildren(h.arr, left, right)
		h.swap(parent, smaller)
		parent = smaller
		left, right = calculateChildren(parent)
	}
}

func calculateChildren(parent int) (int, int) {
	return parent*2 + 1, parent*2 + 2
}

//This function checks if it is feasible to do bubbledown given the left and right indices of the children.
func validBubbleDown[T Number](left, right, parent int, arr []T) bool {
	return left < len(arr) && right < len(arr) && (arr[parent] > arr[left] || arr[parent] > arr[right])
}

func getSmallerChildren[T Number](arr []T, left int, right int) int {
	if (arr)[left] < (arr)[right] {
		return left
	}
	return right
}

//Swap two values given their indexes
func (h *Heap[T]) swap(first int, second int) {
	tmp := h.arr[first]
	h.arr[first] = h.arr[second]
	h.arr[second] = tmp
}

func parent(index int) int {
	return (index - 1) / 2
}
