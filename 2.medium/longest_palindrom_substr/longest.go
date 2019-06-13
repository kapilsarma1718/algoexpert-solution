package lps

func LongestPalindromeSubStr1(str string) string {
	longestString := ""

	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			newWord := str[i:j]

			palindrome := isPalindrome(newWord)
			if palindrome {
				if len(newWord) > len(longestString) {
					longestString = newWord
				}

			}
		}
	}

	return longestString
}

func isPalindrome(str string) bool {
	lenStr := len(str)
	middle := lenStr / 2

	for i := 0; i < middle; i++ {
		left := str[i]
		right := str[lenStr-1-i]

		if left != right {
			return false
		}
	}

	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func LongestPalindromeSubStr2(str string) string {
	return ""
}
