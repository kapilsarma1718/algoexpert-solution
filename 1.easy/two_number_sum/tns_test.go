package twonumbersum_test

import (
	"fmt"
	"reflect"
	"testing"

	td "algoexpert-solution/1.easy/two_number_sum"
)

type TestCase struct {
	arg1 []int
	arg2 int
	want []int
}

func TestTwoNumberSum1(t *testing.T) {
	caseData := []TestCase{
		{
			[]int{4, 6, 1},
			5,
			[]int{1, 4},
		},
		{
			[]int{4, 6, 1, -3},
			3,
			[]int{-3, 6},
		},
		{
			[]int{3, 5, -4, 8, 11, 1, -1, 6},
			15,
			[]int{},
		},
		{
			[]int{3, 5, -4, 8, 11, 1, -1, 6},
			10,
			[]int{-1, 11},
		},
	}

	for _, tc := range caseData {
		t.Run(fmt.Sprintf("TwoNumberSum1(%v, %v)", tc.arg1, tc.arg2), func(t *testing.T) {
			result := td.TwoNumberSum1(tc.arg1, tc.arg2)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("TwoNumberSum1(%v, %v) = %v, want : %v", tc.arg1, tc.arg2, result, tc.want)
			}
		})

		t.Run(fmt.Sprintf("TwoNumberSum2(%v, %v)", tc.arg1, tc.arg2), func(t *testing.T) {
			result := td.TwoNumberSum2(tc.arg1, tc.arg2)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("TwoNumberSum2(%v, %v) = %v, want : %v", tc.arg1, tc.arg2, result, tc.want)
			}
		})
	}
}
