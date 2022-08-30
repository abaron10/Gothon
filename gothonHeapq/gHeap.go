package gothonHeapq

import (
	"github.com/abaron10/Gothon/gothonSlice"
	"golang.org/x/exp/constraints"
)

//left = parent * 2 + 1
//right = parent * 2 + 2

type Number interface {
	constraints.Integer | constraints.Float
}

type Heap[T Number] struct {
	arr  []T
	size int
}

func Heapify[T Number](arr *[]T) {
	heap := Heap[T]{arr: []T{}}
	for _, val := range *arr {
		heap.insert(val)
	}
	*arr = heap.arr
}

func HeapPush(arr *[]int, val int) {
	*arr = append(*arr, val)
	Heapify(arr)
}

func HeapPushPop(arr *[]int, val int) {

}

func HeapPop[T Number](arr *[]T) {
	lastItem := gothonSlice.Pop(arr, -1)
	if len(*arr) == 0 {
		return
	}
	gothonSlice.Remove(arr, (*arr)[0])
	gothonSlice.Insert(arr, 0, lastItem)
	heap := Heap[T]{arr: (*arr)}
	heap.bubbleDown()
	*arr = heap.arr
}

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

//helpers methods
func calculateChildren(parent int) (int, int) {
	return parent*2 + 1, parent*2 + 2
}

func validBubbleDown[T Number](left, right, parent int, arr []T) bool {
	return left < len(arr) && right < len(arr) && (arr[parent] > arr[left] || arr[parent] > arr[right])
}

func getSmallerChildren[T Number](arr []T, left int, right int) int {
	if (arr)[left] < (arr)[right] {
		return left
	}
	return right
}

func (h *Heap[T]) swap(first int, second int) {
	tmp := h.arr[first]
	h.arr[first] = h.arr[second]
	h.arr[second] = tmp
}

func parent(index int) int {
	return (index - 1) / 2
}
