package sd

import "sort"

// O(n)
func SmallestDifference1(arr1, arr2 []int) []int {
	store := make([]int, 2)
	diff := 9999
	tmpDiff := 0

	for i := 0; i < len(arr1); i++ {
		num1 := arr1[i]
		for j := 0; j < len(arr2); j++ {
			num2 := arr2[j]

			if num1 > num2 {
				tmpDiff = num1 - num2
			} else if num2 > num1 {
				tmpDiff = num2 - num1
			} else if num1 == num2 {
				return []int{num1, num2}
			}

			if tmpDiff < diff {
				diff = tmpDiff
				store[0] = num1
				store[1] = num2
			}
		}
	}

	return store

}

func SmallestDifference2(arr1, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	store := make([]int, 2)
	diff := 9999

	left, right := 0, 0

	for left < len(arr1) && right < len(arr2) {
		tmpDiff := 0

		num1 := arr1[left]
		num2 := arr2[right]

		if num1 == num2 {
			return []int{num1, num2}
		} else if num1 > num2 {
			tmpDiff = num1 - num2
			right++
		} else if num2 > num1 {
			tmpDiff = num2 - num1
			left++
		}

		if diff > tmpDiff {
			diff = tmpDiff
			store[0] = num1
			store[1] = num2
		}

	}

	return store
}
