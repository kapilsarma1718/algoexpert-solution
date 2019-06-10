package fibo_test

import (
	"fmt"
	"testing"

	td "algoexpert-solution/1.easy/nth_fibonacci"
)

type testCase struct {
	arg  int
	want int
}

func TestGetNthFib(t *testing.T) {
	caseData := []testCase{
		{
			6,
			5,
		},
		{
			1,
			0,
		},
		{
			2,
			1,
		},
		{
			3,
			1,
		},
		{
			4,
			2,
		},
		{
			5,
			3,
		},
		{
			6,
			5,
		},
	}

	for _, tc := range caseData {
		t.Run(fmt.Sprintf("GetNthFib1(%v)", tc.arg), func(t *testing.T) {
			result := td.GetNthFib1(tc.arg)
			if result != tc.want {
				t.Errorf("GetNthFib1(%v) = %v, want : %v", tc.arg, result, tc.want)
			}
		})

		t.Run(fmt.Sprintf("GetNthFib2(%v)", tc.arg), func(t *testing.T) {
			result := td.GetNthFib2(tc.arg)
			if result != tc.want {
				t.Errorf("GetNthFib2(%v) = %v, want : %v", tc.arg, result, tc.want)
			}
		})
	}

}
