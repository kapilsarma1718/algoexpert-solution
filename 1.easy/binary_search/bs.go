package bs

func BinarySearch(array []int, target int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		middle := (left + right) / 2
		num := array[middle]

		if target == num {
			return middle
		} else if target < num {
			right = middle - 1
		} else if target > num {
			left = middle + 1
		}
	}

	return -1
}
