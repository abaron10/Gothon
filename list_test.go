package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
			ExtendInt([]int{1, 2, 3}, []int{4, 5, 6}),
			[]int{1, 2, 3, 4, 5, 6},
		},
		{"Extends int32",
			ExtendInt32([]int32{1, 2, 3}, []int32{4, 5, 6}),
			[]int32{1, 2, 3, 4, 5, 6},
		},
		{"Extends int32",
			ExtendInt64([]int64{1, 2, 3}, []int64{4, 5, 6}),
			[]int64{1, 2, 3, 4, 5, 6},
		},
		{"Extends float32",
			ExtendFloat32([]float32{1.1, 1.2, 1.3}, []float32{1.4, 1.5, 1.6}),
			[]float32{1.1, 1.2, 1.3, 1.4, 1.5, 1.6},
		},
		{"Extends float64",
			ExtendFloat64([]float64{1.1, 1.2, 1.3}, []float64{1.4, 1.5, 1.6}),
			[]float64{1.1, 1.2, 1.3, 1.4, 1.5, 1.6},
		},
		{"Extends string",
			ExtendString([]string{"apple", "amazon", "google"}, []string{"facebook", "paypal", "netflix"}),
			[]string{"apple", "amazon", "google", "facebook", "paypal", "netflix"},
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
			PopInt(&[]int{1, 2, 3}, -1),
			int(3),
		},
		{"Pop int32",
			PopInt32(&[]int32{1, 2, 3}, 0),
			int32(1),
		},
		{"Pop int64",
			PopInt64(&[]int64{1, 2, 3}, 1),
			int64(2),
		},
		{"Pop float32",
			PopFloat32(&[]float32{1.1, 1.2, 1.3}, 2),
			float32(1.3),
		},
		{"Pop float64",
			PopFloat64(&[]float64{1.1, 1.2, 1.3}, -2),
			1.2,
		},
		{"Pop string",
			PopString(&[]string{"apple", "amazon", "google"}, -1),
			"google",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PoppedValue := tt.extendFunction
			assert.Equal(t, tt.expectedPoppedValue, PoppedValue)
		})
	}
}
