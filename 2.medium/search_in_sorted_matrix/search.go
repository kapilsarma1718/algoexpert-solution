package matrixsearch

// O(n * m)
func SearchInSortedMatrix1(matrix [][]int, target int) []int {

	for row := 0; row < len(matrix); row++ {
		lenRow := len(matrix[row])
		for col := 0; col < lenRow; col++ {
			num := matrix[row][col]

			if num == target {
				return []int{row, col}
			}
		}
	}

	return []int{-1, 1}
}

// O(n + m)
func SearchInSortedMatrix2(matrix [][]int, target int) []int {

	row := 0
	col := len(matrix[row]) - 1

	for row < len(matrix) && col >= 0 {
		num := matrix[row][col]

		if target == num {
			return []int{row, col}
		} else if target < num {
			col--
		} else if target > num {
			row++
		}
	}

	return []int{-1, 1}

}
