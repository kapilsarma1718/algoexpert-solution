package main

import (
	"fmt"
	"reflect"
	"testing"
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
			[]int{3, -6},
		},
		{
			[]int{3, 5, -4, 8, 11, 1, -1, 6},
			15,
			[]int{},
		},
	}

	for _, tc := range caseData {
		t.Run(fmt.Sprintf("TwoNumberSum1(%v, %v)", tc.arg1, tc.arg2), func(t *testing.T) {
			result := TwoNumberSum1(tc.arg1, tc.arg2)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("TwoNumberSum1(%v, %v) = %v, want : %v", tc.arg1, tc.arg2, result, tc.want)
			}
		})
	}
}
