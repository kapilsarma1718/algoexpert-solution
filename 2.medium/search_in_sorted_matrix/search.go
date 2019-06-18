package matrixsearch

// O(n ^ 2)
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
