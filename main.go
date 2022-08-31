package main

import (
	"fmt"
	"github.com/abaron10/Gothon/gothonHeapq"
)

func main() {
	s := NewMedian()

	for i := 0; i < 10000; i++ {
		s.AddNumber(i)
	}

	fmt.Println(s.FindMedian())
}

type MedianFinder struct {
	maxHeap *[]int
	minHeap *[]int
}

func NewMedian() *MedianFinder {
	return &MedianFinder{&[]int{}, &[]int{}}
}

func (self *MedianFinder) AddNumber(num int) {
	gothonHeapq.HeapPush(self.maxHeap, -num)
	gothonHeapq.HeapPush(self.minHeap, -(*self.maxHeap)[0])
	gothonHeapq.HeapPop(self.maxHeap)

	if len(*self.minHeap) > len(*self.maxHeap) {
		gothonHeapq.HeapPush(self.maxHeap, -(*self.minHeap)[0])
		gothonHeapq.HeapPop(self.minHeap)
	}
}

func (self *MedianFinder) FindMedian() float64 {
	if len(*self.maxHeap) == len(*self.minHeap) {
		return float64(-(*self.maxHeap)[0]+(*self.minHeap)[0]) / 2
	} else {
		return float64(-(*self.maxHeap)[0])
	}
}
