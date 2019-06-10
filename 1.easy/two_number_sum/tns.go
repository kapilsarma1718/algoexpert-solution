package twonumbersum

func TwoNumberSum1(array []int, target int) []int {

	result := []int{}
	lenArr := len(array)

	for i := 0; i < lenArr-1; i++ {
		for j := i + 1; j < lenArr; j++ {
			num1 := array[i]
			num2 := array[j]

			total := num1 + num2

			if total == target {
				minValue := Min(num1, num2)
				maxValue := Max(num1, num2)
				result = append(result, minValue, maxValue)
				return result
			}
		}
	}

	return []int{}
}

func TwoNumberSum2(array []int, target int) []int {
	mapValue := map[int]int{}
	result := []int{}

	for _, num := range array {
		n2 := target - num

		value, ok := mapValue[n2]
		if ok {
			minValue := Min(n2, value)
			maxValue := Max(n2, value)

			result = append(result, minValue, maxValue)
			return result
		}

		mapValue[num] = n2
	}

	return result

}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
