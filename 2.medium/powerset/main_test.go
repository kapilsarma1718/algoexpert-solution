package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	arg  []int
	want [][]int
}

func TestPowerset(t *testing.T) {

	caseData := []testCase{
		{
			[]int{},
			[][]int{
				[]int{},
			},
		},
		{
			[]int{1, 2, 3},
			[][]int{
				[]int{},
				[]int{1},
				[]int{2},
				[]int{1, 2},
				[]int{3},
				[]int{1, 2},
				[]int{2, 3},
				[]int{1, 2, 3},
			},
		},
	}

	for _, tc := range caseData {
		t.Run(fmt.Sprintf("Powerset(%v)", tc.arg), func(t *testing.T) {

			result := Powerset(tc.arg)
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("Powerset(%v) = %v, want : %v", tc.arg, result, tc.want)
			}

		})
	}

}
