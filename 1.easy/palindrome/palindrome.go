package palindrome

func IsPalindrome1(str string) bool {

	if str == "" {
		return false
	}

	tmpStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		tmpStr = tmpStr + string(str[i])
	}

	return str == tmpStr
}

func IsPalindrome2(str string) bool {

	if str == "" {
		return false
	}

	mid := len(str) / 2

	for i := 0; i < mid; i++ {
		left := str[i]
		right := str[len(str)-1-i]

		if left != right {
			return false
		}
	}

	return true

}
