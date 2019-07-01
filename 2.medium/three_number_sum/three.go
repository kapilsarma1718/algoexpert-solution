package tns

import (
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {
	store := [][]int{}
	sort.Ints(array)

	for i := 0; i < len(array)-2; i++ {
		num := array[i]
		left := i + 1
		right := len(array) - 1

		for left < right {
			res := num + array[left] + array[right]

			if res == target {
				store = append(store, []int{num, array[left], array[right]})
				left++
				right--
			} else if res > target {
				right--
			} else {
				left++
			}
		}

	}

	return store
}
