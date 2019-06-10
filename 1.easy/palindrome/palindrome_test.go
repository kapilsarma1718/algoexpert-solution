package palindrome_test

import (
	"fmt"
	"testing"

	td "algoexpert-solution/1.easy/palindrome"
)

type testCase struct {
	arg  string
	want bool
}

func TestIsPalindrome(t *testing.T) {
	caseData := []testCase{
		{
			"bob",
			true,
		},
		{
			"abcdcba",
			true,
		},
		{
			"dfasdf",
			false,
		},
		{
			"",
			false,
		},
	}

	for _, tc := range caseData {
		t.Run(fmt.Sprintf("IsPalindrome1(%v)", tc.arg), func(t *testing.T) {
			result := td.IsPalindrome1(tc.arg)
			if result != tc.want {
				t.Errorf("IsPalindrome1(%v) = %v, want : %v", tc.arg, result, tc.want)
			}
		})

		t.Run(fmt.Sprintf("IsPalindrome2(%v)", tc.arg), func(t *testing.T) {
			result := td.IsPalindrome2(tc.arg)
			if result != tc.want {
				t.Errorf("IsPalindrome1(%v) = %v, want : %v", tc.arg, result, tc.want)
			}
		})
	}
}
