package ftln

import (
	"math"
)

func FindThreeLargestNumbers(array []int) []int {
	result := []int{math.MinInt32, math.MinInt32, math.MinInt32}

	for i := 0; i < len(array); i++ {
		num := array[i]

		if num > result[2] {
			result[0] = result[1]
			result[1] = result[2]
			result[2] = num
		} else if num > result[1] {
			result[0] = result[1]
			result[1] = num
		} else if num > result[0] {
			result[0] = num
		}
	}

	return result

}
