package fibo

func GetNthFib1(n int) int {
	result := 0
	nums := []int{0, 1}

	if n <= 2 {
		return nums[n-1]
	}

	for i := 2; i < n; i++ {
		result = nums[i-2] + nums[i-1]
		nums = append(nums, result)
	}

	return result
}

func GetNthFib2(n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	return GetNthFib2(n-1) + GetNthFib2(n-2)
}
