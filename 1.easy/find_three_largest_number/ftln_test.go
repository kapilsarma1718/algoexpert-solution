package ftln_test

import (
	"fmt"
	"reflect"
	"testing"

	td "algoexpert-solution/1.easy/find_three_largest_number"
)

type testCase struct {
	arg  []int
	want []int
}

func TestFindThreeLargestNumbers(t *testing.T) {

	caseData := []testCase{
		{
			[]int{55, 7, 8},
			[]int{7, 8, 55},
		},
		{
			[]int{-1, -2, -3, -7, -17, -27, -18, -541, -8, -7, 7},
			[]int{-2, -1, 7},
		},
	}

	for _, tc := range caseData {
		t.Run(fmt.Sprintf("FindThreeLargestNumbers(%v)", tc.arg), func(t *testing.T) {
			result := td.FindThreeLargestNumbers(tc.arg)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("FindThreeLargestNumbers(%v) = %v, want : %v", tc.arg, result, tc.want)
			}
		})
	}

}
