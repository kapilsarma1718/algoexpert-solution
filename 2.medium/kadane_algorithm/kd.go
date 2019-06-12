package kd

// BruteForce - O(n^2)
func BruteForce(array []int) int {
	maxSum := array[0]

	for i := 0; i < len(array); i++ {

		// every value, tmpSum set to 0
		// because it tmpSum will be add with contigous subarray from index i to end
		tmpSum := 0
		for j := i; j < len(array); j++ {
			num := array[j]
			tmpSum = tmpSum + num

			if tmpSum > maxSum {
				maxSum = tmpSum
			}
		}
	}

	return maxSum
}

func Kadanes(array []int) int {
	return 1
}
