package bs_test

import (
	"fmt"
	"reflect"
	"testing"

	tt "algoexpert-solution/1.easy/bubble_sort"
)

type TestCase struct {
	arg  []int
	want []int
}

func TestBinarySearch(t *testing.T) {
	caseData := []TestCase{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1, 3, 2},
			[]int{1, 2, 3},
		},
		{
			[]int{-4, 5, 10, 8, -10, -6, -4, -2, -5, 3, 5, -4, -5, -1, 1, 6, -7, -6, -7, 8},
			[]int{-10, -7, -7, -6, -6, -5, -5, -4, -4, -4, -2, -1, 1, 3, 5, 5, 6, 8, 8, 10},
		},
	}

	for _, tc := range caseData {
		t.Run(fmt.Sprintf("BubbleSort(%v)", tc.arg), func(t *testing.T) {
			result := tt.BubbleSort(tc.arg)

			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("BubbleSort(%v) = %v, want : %v", tc.arg, result, tc.want)
			}

		})
	}

}
