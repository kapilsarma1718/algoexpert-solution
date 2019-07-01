package bs

func BubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			num1 := array[j]
			num2 := array[j+1]

			if num1 > num2 {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
