package lps_test

import (
	"fmt"
	"testing"

	td "algoexpert-solution/2.medium/longest_palindrom_substr"
)

type testCase struct {
	arg  string
	want string
}

func TestLongestPalindromeSubStr(t *testing.T) {
	caseData := []testCase{
		{
			"a",
			"a",
		},
		{
			"abcdefgfedcbazzzzzzzzzzzzzzzzzzzz",
			"zzzzzzzzzzzzzzzzzzzz",
		},
		{
			"abcdefghfedcbaa",
			"aa",
		},
	}

	for _, tc := range caseData {
		t.Run(fmt.Sprintf("LongestPalindromeSubStr1(%v)", tc.arg), func(t *testing.T) {
			result := td.LongestPalindromeSubStr1(tc.arg)

			if result != tc.want {
				t.Errorf("LongestPalindromeSubStr1(%v) = %v, want : %v", tc.arg, result, tc.want)
			}

		})
	}

}
