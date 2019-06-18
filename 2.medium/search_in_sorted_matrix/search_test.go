package matrixsearch_test

import (
	"fmt"
	"reflect"
	"testing"

	td "algoexpert-solution/2.medium/search_in_sorted_matrix"
)

type testCase struct {
	arg1 [][]int
	arg2 int
	want []int
}

func TestSearchInSortedMatrix(t *testing.T) {
	caseData := []testCase{
		{
			[][]int{
				{1, 4, 7, 12, 15, 1000},
				{2, 5, 19, 31, 32, 1001},
				{3, 8, 24, 33, 35, 1002},
				{40, 41, 42, 44, 45, 1003},
				{99, 100, 103, 106, 128, 1004},
			},
			44,
			[]int{3, 3},
		},
		{
			[][]int{
				{1, 2},
				{4, 5},
			},
			10,
			[]int{-1, 1},
		},
	}

	for _, tc := range caseData {
		t.Run(fmt.Sprintf("SearchInSortedMatrix1()"), func(t *testing.T) {
			result := td.SearchInSortedMatrix1(tc.arg1, tc.arg2)

			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("SearchInSortedMatrix1(%v, %v) = %v, want : %v", tc.arg1, tc.arg2, result, tc.want)
			}

		})

		t.Run(fmt.Sprintf("SearchInSortedMatrix2()"), func(t *testing.T) {
			result := td.SearchInSortedMatrix2(tc.arg1, tc.arg2)

			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("SearchInSortedMatrix2(%v, %v) = %v, want : %v", tc.arg1, tc.arg2, result, tc.want)
			}

		})
	}
}
