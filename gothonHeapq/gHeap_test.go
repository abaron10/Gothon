package gothonHeapq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------HEAPIFY METHOD TESTS---------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/
func TestHeapifyMethod(t *testing.T) {
	t.Run("Heapify slice of int", func(t *testing.T) {
		a := []int{100, 1, 3, 300, 5}
		Heapify(&a)
		minimum := a[0]
		assert.Equal(t, minimum, 1)
	})
	t.Run("Heapify slice of int32", func(t *testing.T) {
		a := []int32{100, -31, 3, 300, 5}
		Heapify(&a)
		minimum := a[0]
		assert.Equal(t, minimum, int32(-31))
	})
	t.Run("Heapify slice of int64", func(t *testing.T) {
		a := []int64{100, 1, 3, 300, 5}
		Heapify(&a)
		minimum := a[0]
		assert.Equal(t, minimum, int64(1))
	})
	t.Run("Heapify slice of float32", func(t *testing.T) {
		a := []float32{100.0, 1.0, 3.0, 0.1, 5.0}
		Heapify(&a)
		minimum := a[0]
		assert.Equal(t, minimum, float32(0.1))
	})
	t.Run("Heapify slice of float64", func(t *testing.T) {
		a := []float64{100.0, 1.0, 3.0, 0.1, 5.0}
		Heapify(&a)
		minimum := a[0]
		assert.Equal(t, minimum, float64(0.1))
	})
}

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------HEAPPUSH METHOD TESTS---------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/
func TestHeapPushMethod(t *testing.T) {
	t.Run("HeapPush slice of int", func(t *testing.T) {
		a := []int{100, 1, 3, 300, 5}
		HeapPush(&a, -30)
		minimum := a[0]
		assert.Equal(t, minimum, -30)
	})
	t.Run("HeapPush slice of int32", func(t *testing.T) {
		a := []int32{100, -31, 3, 300, 5}
		HeapPush(&a, -60)
		minimum := a[0]
		assert.Equal(t, minimum, int32(-60))
	})
	t.Run("HeapPush slice of int64", func(t *testing.T) {
		a := []int64{100, 1, 3, 300, 5}
		HeapPush(&a, -90)
		minimum := a[0]
		assert.Equal(t, minimum, int64(-90))
	})
	t.Run("HeapPush slice of float32", func(t *testing.T) {
		a := []float32{100.0, 1.0, 3.0, 0.1, 5.0}
		HeapPush(&a, -120)
		minimum := a[0]
		assert.Equal(t, minimum, float32(-120))
	})
	t.Run("HeapPush slice of float64", func(t *testing.T) {
		a := []float64{100.0, 1.0, 3.0, 0.1, 5.0}
		HeapPush(&a, -150)
		minimum := a[0]
		assert.Equal(t, minimum, float64(-150))
	})
}

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------HEAPPUSHPOP METHOD TESTS----------------------------------------
-----------------------------------------------------------------------------------------------------------
*/
func TestHeapPushPopMethod(t *testing.T) {
	t.Run("HeapPushPop slice of int", func(t *testing.T) {
		a := []int{100, 1, 3, 300, 5}
		minimumPopped := HeapPushPop(&a, -30)
		CurrentMinimum := a[0]
		assert.Equal(t, CurrentMinimum, 1)
		assert.Equal(t, minimumPopped, -30)
	})
	t.Run("HeapPushPop slice of int32", func(t *testing.T) {
		a := []int32{100, -31, 3, 300, 5}
		minimumPopped := HeapPushPop(&a, -30)
		CurrentMinimum := a[0]
		assert.Equal(t, CurrentMinimum, int32(-30))
		assert.Equal(t, minimumPopped, int32(-31))
	})
	t.Run("HeapPushPop slice of int64", func(t *testing.T) {
		a := []int64{100, 1, 3, 300, 5}
		minimumPopped := HeapPushPop(&a, -30)
		CurrentMinimum := a[0]
		assert.Equal(t, CurrentMinimum, int64(1))
		assert.Equal(t, minimumPopped, int64(-30))
	})
	t.Run("HeapPushPop slice of float32", func(t *testing.T) {
		a := []float32{100.0, 1.0, 3.0, 0.1, 5.0}
		minimumPopped := HeapPushPop(&a, -30)
		CurrentMinimum := a[0]
		assert.Equal(t, CurrentMinimum, float32(0.1))
		assert.Equal(t, minimumPopped, float32(-30))
	})
	t.Run("HeapPushPop slice of float64", func(t *testing.T) {
		a := []float64{100.0, 1.0, 3.0, 0.1, 5.0}
		minimumPopped := HeapPushPop(&a, -30)
		CurrentMinimum := a[0]
		assert.Equal(t, CurrentMinimum, float64(0.1))
		assert.Equal(t, minimumPopped, float64(-30))
	})
}

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------HEAPPOP METHOD TESTS--------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/
func TestHeapPopMethod(t *testing.T) {
	t.Run("HeapPop slice of int", func(t *testing.T) {
		a := []int{100, 1, 3, 300, 5}
		Heapify(&a)
		minimumPopped := HeapPop(&a)
		assert.Equal(t, minimumPopped, 1)
	})
	t.Run("HeapPop slice of int32", func(t *testing.T) {
		a := []int32{100, -31, 3, 300, 5}
		Heapify(&a)
		minimumPopped := HeapPop(&a)
		assert.Equal(t, minimumPopped, int32(-31))
	})
	t.Run("HeapPop slice of int64", func(t *testing.T) {
		a := []int64{100, 1, 3, 300, 5}
		Heapify(&a)
		minimumPopped := HeapPop(&a)
		assert.Equal(t, minimumPopped, int64(1))
	})
	t.Run("HeapPop slice of float32", func(t *testing.T) {
		a := []float32{100.0, 1.0, 3.0, 0.1, 5.0}
		Heapify(&a)
		minimumPopped := HeapPop(&a)
		assert.Equal(t, minimumPopped, float32(0.1))
	})
	t.Run("HeapPop slice of float64", func(t *testing.T) {
		a := []float64{100.0, 1.0, 3.0, 0.1, 5.0}
		Heapify(&a)
		minimumPopped := HeapPop(&a)
		assert.Equal(t, minimumPopped, float64(0.1))
	})
}
