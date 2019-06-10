package bs_test

import (
	"testing"

	td "algoexpert-solution/1.easy/binary_search"
)

func TestBinarySearch(t *testing.T) {
	arg := []int{10, 14, 19, 26, 27, 31, 33, 35, 42, 44}
	want := 5

	result := td.BinarySearch(arg, 31)
	if result != want {
		t.Errorf("BinarySearch(%v) = %v, want : %v", arg, result, want)
	}

}
