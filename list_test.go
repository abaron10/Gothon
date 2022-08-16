package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*---------------------------------------------------------------------------------------------------------
-------------------------------------------EXTEND METHOD TESTS---------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/

type TestStruct struct {
	id int
}

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
		extendFunction      interface{}
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
			PoppedValue := tt.extendFunction
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
		extendFunction     interface{}
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
			IndexValue := tt.extendFunction
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
		name               string
		extendFunction     interface{}
		expectedIndexValue interface{}
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
			IndexValue := tt.extendFunction
			assert.Equal(t, tt.expectedIndexValue, IndexValue)
		})
	}
}

/*---------------------------------------------------------------------------------------------------------
-----------------------------------------REVERSE METHOD TESTS----------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/

func TestReverseMethod(t *testing.T) {
	var tests = []struct {
		name           string
		sliceToReverse interface{}
		reversedSlice  interface{}
	}{
		{"Reverse int",
			[]int{1, 2, 3},
			[]int{3, 2, 1},
		},
		{"Reverse int32",
			[]int32{1, 4, 9},
			[]int32{9, 4, 1},
		},
		{"Reverse int64",
			[]int64{9, 8, 7},
			[]int64{7, 8, 9},
		},
		{"Reverse float32",
			[]float32{1.1, 1.2, 1.3},
			[]float32{1.3, 1.2, 1.1},
		},
		{"Reverse float64",
			[]float64{9.3, 9.1, 9.4},
			[]float64{9.4, 9.1, 9.3},
		},
		{"Reverse string",
			[]string{"apple", "amazon", "google"},
			[]string{"google", "amazon", "apple"},
		},
		{"Reverse struct",
			[]TestStruct{{1}, {2}},
			[]TestStruct{{2}, {1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			Reverse(&tt.sliceToReverse)
			assert.Equal(t, tt.expectedIndexValue, IndexValue)
		})
	}
}
