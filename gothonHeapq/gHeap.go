package gothonHeapq

type T any

type heap struct{}

var heapC []int

func Heapify(a *[]int) {
	for _, item := range *a {
		Heap(item)
	}
	(*a) = heapC
}

func HeapPush(a *[]int, value int) {
	*a = append(*a, value)
	Heapify(a)
}

func Heap(value int) {
	prev := len(heapC) - 1
	if prev < 0 || heapC[prev] < value {
		heapC = append(heapC, value)
		return
	}
	heapC = append(heapC, 0)
	for prev > -1 && value < heapC[prev] {
		heapC[prev+1] = heapC[prev]
		heapC[prev] = value
		prev--
	}
}
