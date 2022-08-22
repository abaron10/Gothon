package gothonSlice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	id int
}

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------EXTEND METHOD TESTS---------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/
func TestExtendMethod(t *testing.T) {

	var tests = []struct {
		name           string
		extendFunction interface{}
		expectedResult interface{}
	}{
		{"Extends int",
			Extend([]int{1, 2, 3}, []int{4, 5, 6}),
			[]int{1, 2, 3, 4, 5, 6},
		},
		{"Extends int32",
			Extend([]int32{1, 2, 3}, []int32{4, 5, 6}),
			[]int32{1, 2, 3, 4, 5, 6},
		},
		{"Extends int32",
			Extend([]int64{1, 2, 3}, []int64{4, 5, 6}),
			[]int64{1, 2, 3, 4, 5, 6},
		},
		{"Extends float32",
			Extend([]float32{1.1, 1.2, 1.3}, []float32{1.4, 1.5, 1.6}),
			[]float32{1.1, 1.2, 1.3, 1.4, 1.5, 1.6},
		},
		{"Extends float64",
			Extend([]float64{1.1, 1.2, 1.3}, []float64{1.4, 1.5, 1.6}),
			[]float64{1.1, 1.2, 1.3, 1.4, 1.5, 1.6},
		},
		{"Extends string",
			Extend([]string{"apple", "amazon", "google"}, []string{"facebook", "paypal", "netflix"}),
			[]string{"apple", "amazon", "google", "facebook", "paypal", "netflix"},
		},
		{"Extends struct",
			Extend([]TestStruct{{1}}, []TestStruct{{2}}),
			[]TestStruct{{1}, {2}},
		},
		{"Extends string first slice empty and second with values",
			Extend([]string{}, []string{"facebook", "paypal", "netflix"}),
			[]string{"facebook", "paypal", "netflix"},
		},
		{"Extends string first slice with values and second empty",
			Extend([]string{"apple", "amazon", "google"}, []string{}),
			[]string{"apple", "amazon", "google"},
		},
		{"Extends string first and second slice empty",
			Extend([]string{}, []string{}),
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extendedList := tt.extendFunction
			assert.Equal(t, tt.expectedResult, extendedList)
		})
	}
}

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------POP METHOD TESTS------------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/

func TestPopMethod(t *testing.T) {
	var tests = []struct {
		name                string
		popFunction         interface{}
		expectedPoppedValue interface{}
	}{
		{"Pop int",
			Pop(&[]int{1, 2, 3}, -1),
			int(3),
		},
		{"Pop int32",
			Pop(&[]int32{1, 2, 3}, 0),
			int32(1),
		},
		{"Pop int64",
			Pop(&[]int64{1, 2, 3}, 1),
			int64(2),
		},
		{"Pop float32",
			Pop(&[]float32{1.1, 1.2, 1.3}, 2),
			float32(1.3),
		},
		{"Pop float64",
			Pop(&[]float64{1.1, 1.2, 1.3}, -2),
			1.2,
		},
		{"Pop string",
			Pop(&[]string{"apple", "amazon", "google"}, -1),
			"google",
		},
		{"Pop struct",
			Pop(&[]TestStruct{{1}, {2}}, -1),
			TestStruct{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PoppedValue := tt.popFunction
			assert.Equal(t, tt.expectedPoppedValue, PoppedValue)
		})
	}
}

/*---------------------------------------------------------------------------------------------------------
-----------------------------------------INDEX METHOD TESTS------------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/

func TestIndexMethod(t *testing.T) {
	var tests = []struct {
		name               string
		indexFunction      interface{}
		expectedIndexValue interface{}
	}{
		{"Index int",
			Index([]int{1, 2, 3}, -1),
			-1,
		},
		{"Index int32",
			Index([]int32{1, 2, 3}, 0),
			-1,
		},
		{"Index int64",
			Index([]int64{1, 2, 3}, 1),
			0,
		},
		{"Index float32",
			Index([]float32{1.1, 1.2, 1.3}, 1.3),
			2,
		},
		{"Index float64",
			Index([]float64{1.1, 1.2, 1.3}, 1.2),
			1,
		},
		{"Index string",
			Index([]string{"apple", "amazon", "google"}, "google"),
			2,
		},
		{"Index struct",
			Index([]TestStruct{{1}, {2}}, TestStruct{1}),
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IndexValue := tt.indexFunction
			assert.Equal(t, tt.expectedIndexValue, IndexValue)
		})
	}
}

/*---------------------------------------------------------------------------------------------------------
----------------------------------------CONTAINS METHOD TESTS----------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/

func TestContainsMethod(t *testing.T) {
	var tests = []struct {
		name                     string
		containsFunction         interface{}
		expectedContainsResponse interface{}
	}{
		{"Contains int",
			Contains([]int{1, 2, 3}, -1),
			false,
		},
		{"Contains int32",
			Contains([]int32{1, 2, 3}, 0),
			false,
		},
		{"Contains int64",
			Contains([]int64{1, 2, 3}, 1),
			true,
		},
		{"Contains float32",
			Contains([]float32{1.1, 1.2, 1.3}, 1.3),
			true,
		},
		{"Contains float64",
			Contains([]float64{1.1, 1.2, 1.3}, 1.2),
			true,
		},
		{"Contains string",
			Contains([]string{"apple", "amazon", "google"}, "google"),
			true,
		},
		{"Contains struct",
			Contains([]TestStruct{{1}, {2}}, TestStruct{1}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IndexValue := tt.containsFunction
			assert.Equal(t, tt.expectedContainsResponse, IndexValue)
		})
	}
}

/*---------------------------------------------------------------------------------------------------------
-----------------------------------------REVERSE METHOD TESTS----------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/
func TestReverseMethod(t *testing.T) {

	t.Run("Reverse int", func(t *testing.T) {
		a := []int{1, 2, 3}
		Reverse(&a)
		assert.Equal(t, []int{3, 2, 1}, a)
	})
	t.Run("Reverse int32", func(t *testing.T) {
		a := []int32{1, 2, 3}
		Reverse(&a)
		assert.Equal(t, []int32{3, 2, 1}, a)
	})
	t.Run("Reverse int64", func(t *testing.T) {
		a := []int64{1, 2, 3}
		Reverse(&a)
		assert.Equal(t, []int64{3, 2, 1}, a)
	})
	t.Run("Reverse float32", func(t *testing.T) {
		a := []float32{1.1, 2.1, 3.1}
		Reverse(&a)
		assert.Equal(t, []float32{3.1, 2.1, 1.1}, a)
	})
	t.Run("Reverse float64", func(t *testing.T) {
		a := []float64{1.1, 2.1, 3.1}
		Reverse(&a)
		assert.Equal(t, []float64{3.1, 2.1, 1.1}, a)
	})
	t.Run("Reverse string", func(t *testing.T) {
		a := []string{"apple", "amazon", "google"}
		Reverse(&a)
		assert.Equal(t, []string{"google", "amazon", "apple"}, a)
	})
	t.Run("Reverse struct", func(t *testing.T) {
		a := []TestStruct{{1}, {2}}
		Reverse(&a)
		assert.Equal(t, []TestStruct{{2}, {1}}, a)
	})
}

/*---------------------------------------------------------------------------------------------------------
------------------------------------------REMOVE METHOD TESTS----------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/
func TestRemoveMethod(t *testing.T) {

	t.Run("Remove int", func(t *testing.T) {
		a := []int{1, 2, 3}
		Remove(&a, 2)
		assert.Equal(t, []int{1, 3}, a)
	})
	t.Run("Remove int32", func(t *testing.T) {
		a := []int32{1, 2, 3}
		Remove(&a, 1)
		assert.Equal(t, []int32{2, 3}, a)
	})
	t.Run("Remove int64", func(t *testing.T) {
		a := []int64{1, 2, 3}
		Remove(&a, 3)
		assert.Equal(t, []int64{1, 2}, a)
	})
	t.Run("Remove float32", func(t *testing.T) {
		a := []float32{1.1, 2.1, 3.1}
		Remove(&a, 2.1)
		assert.Equal(t, []float32{1.1, 3.1}, a)
	})
	t.Run("Remove float64", func(t *testing.T) {
		a := []float64{1.1, 2.1, 3.1}
		Remove(&a, 1.1)
		assert.Equal(t, []float64{2.1, 3.1}, a)
	})
	t.Run("Remove string", func(t *testing.T) {
		a := []string{"apple", "amazon", "google"}
		Remove(&a, "apple")
		assert.Equal(t, []string{"amazon", "google"}, a)
	})
	t.Run("Remove struct", func(t *testing.T) {
		a := []TestStruct{{1}, {2}}
		Remove(&a, a[1])
		assert.Equal(t, []TestStruct{{1}}, a)
	})
}

/*---------------------------------------------------------------------------------------------------------
------------------------------------------INSERT METHOD TESTS----------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/
func TestInsertMethod(t *testing.T) {

	t.Run("Insert int", func(t *testing.T) {
		a := []int{1, 2, 3}
		Insert(&a, 0, 9)
		assert.Equal(t, []int{9, 1, 2, 3}, a)
	})
	t.Run("Insert int32", func(t *testing.T) {
		a := []int32{1, 2, 3}
		Insert(&a, 1, 9)
		assert.Equal(t, []int32{1, 9, 2, 3}, a)
	})
	t.Run("Insert int64", func(t *testing.T) {
		a := []int64{1, 2, 3}
		Insert(&a, 2, 9)
		assert.Equal(t, []int64{1, 2, 9, 3}, a)
	})
	t.Run("Insert float32", func(t *testing.T) {
		a := []float32{1.1, 2.1, 3.1}
		Insert(&a, 2, 9)
		assert.Equal(t, []float32{1.1, 2.1, 9, 3.1}, a)
	})
	t.Run("Insert float64", func(t *testing.T) {
		a := []float64{1.1, 2.1, 3.1}
		Insert(&a, 0, 1)
		assert.Equal(t, []float64{1, 1.1, 2.1, 3.1}, a)
	})
	t.Run("Insert string", func(t *testing.T) {
		a := []string{"apple", "amazon", "google"}
		Insert(&a, 1, "meta")
		assert.Equal(t, []string{"apple", "meta", "amazon", "google"}, a)
	})
	t.Run("Insert struct", func(t *testing.T) {
		a := []TestStruct{{1}, {2}}
		Insert(&a, 1, a[1])
		assert.Equal(t, []TestStruct{{1}, {2}, {2}}, a)
	})
}

/*---------------------------------------------------------------------------------------------------------
------------------------------------------COPY METHOD TESTS----------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/
func TestCopyMethod(t *testing.T) {

	t.Run("Copy int", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := Copy(a)
		assert.Equal(t, b, a)
	})
	t.Run("Copy int32", func(t *testing.T) {
		a := []int32{1, 2, 3}
		b := Copy(a)
		assert.Equal(t, b, a)
	})
	t.Run("Copy int64", func(t *testing.T) {
		a := []int64{1, 2, 3}
		b := Copy(a)
		assert.Equal(t, b, a)
	})
	t.Run("Copy float32", func(t *testing.T) {
		a := []float32{1.1, 2.1, 3.1}
		b := Copy(a)
		assert.Equal(t, b, a)
	})
	t.Run("Copy float64", func(t *testing.T) {
		a := []float64{1.1, 2.1, 3.1}
		b := Copy(a)
		assert.Equal(t, b, a)
	})
	t.Run("Copy string", func(t *testing.T) {
		a := []string{"apple", "amazon", "google"}
		b := Copy(a)
		assert.Equal(t, b, a)
	})
	t.Run("Copy struct", func(t *testing.T) {
		a := []TestStruct{{1}, {2}}
		b := Copy(a)
		assert.Equal(t, b, a)
	})
}
