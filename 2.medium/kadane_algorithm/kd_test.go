package kd_test

import (
	"fmt"
	"testing"

	td "algoexpert-solution/2.medium/kadane_algorithm"
)

type testCase struct {
	arg  []int
	want int
}

func TestLargestSumContigousSubArray(t *testing.T) {
	caseData := []testCase{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			55,
		},
		{
			[]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			-1,
		},
		{
			[]int{3, 4, -6, 7, 8},
			16,
		},
		{
			[]int{-3, -2, 100},
			100,
		},
		{
			[]int{8, 5, -9, 1, 3, -2, 3, 4, 7, 2, -18, 6, 3, 1, -5, 6, 20, -23, 15, 1, -3, 4},
			35,
		},
	}

	for _, tc := range caseData {
		t.Run(fmt.Sprintf("BruteForce(%v)", tc.arg), func(t *testing.T) {
			result := td.BruteForce(tc.arg)
			if result != tc.want {

				t.Errorf("BruteForce(%v) = %v, want : %v", tc.arg, result, tc.want)
			}
		})
	}
}
