package sd_test

import (
	"fmt"
	"reflect"
	"testing"

	td "algoexpert-solution/2.medium/smallest_difference"
)

type testCase struct {
	arg1 []int
	arg2 []int
	want []int
}

func TestSmallestDifference(t *testing.T) {
	caseData := []testCase{
		{
			[]int{-1, 5, 10, 20, 3},
			[]int{26, 134, 135, 15, 17},
			[]int{20, 17},
		},
		{
			[]int{-1, 5, 10, 20, 28, 3},
			[]int{26, 134, 135, 15, 17},
			[]int{28, 26},
		},
	}

	for _, tc := range caseData {
		t.Run(fmt.Sprintf("SmallestDifference1(%v, %v)", tc.arg1, tc.arg2), func(t *testing.T) {
			result := td.SmallestDifference1(tc.arg1, tc.arg2)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("SmallestDifference1(%v, %v) = %v, want : %v", tc.arg1, tc.arg2, result, tc.want)
			}
		})

		t.Run(fmt.Sprintf("SmallestDifference2(%v, %v)", tc.arg1, tc.arg2), func(t *testing.T) {
			result := td.SmallestDifference2(tc.arg1, tc.arg2)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("SmallestDifference2(%v, %v) = %v, want : %v", tc.arg1, tc.arg2, result, tc.want)
			}
		})
	}
}
