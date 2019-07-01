package tns_test

import (
	"fmt"
	"reflect"
	"testing"

	tt "algoexpert-solution/2.medium/three_number_sum"
)

type testCase struct {
	arg1 []int
	arg2 int
	want [][]int
}

func TestThreeNumberSum(t *testing.T) {
	caseData := []testCase{
		{
			[]int{},
			0,
			[][]int{},
		},
		{
			[]int{12, 3, 1, 2, -6, 5, -8, 6},
			0,
			[][]int{
				[]int{-8, 2, 6},
				[]int{-8, 3, 5},
				[]int{-6, 1, 5},
			},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}, 5,
			[][]int{},
		},
	}

	for _, tc := range caseData {
		t.Run(fmt.Sprintf("ThreeNumberSum(%v, %v)", tc.arg1, tc.arg2), func(t *testing.T) {
			result := tt.ThreeNumberSum(tc.arg1, tc.arg2)

			if !reflect.DeepEqual(result, tc.want) {

				t.Errorf("ThreeNumberSum(%v, %v) = %v, want : %v", tc.arg1, tc.arg2, result, tc.want)
			}

		})
	}
}
