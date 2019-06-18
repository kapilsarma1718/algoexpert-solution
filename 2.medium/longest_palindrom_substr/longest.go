package lps

// O(n ^ 3)
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

// O (n ^ 2)
func LongestPalindromeSubStr2(str string) string {
	longest := str[0:1]

	for i := 1; i < len(str); i++ {
		tmpLongest := ""
		odd := getLongestPalindrome(str, i-1, i+1)
		even := getLongestPalindrome(str, i-1, i)

		if len(odd) > len(even) {
			tmpLongest = odd
		} else {
			tmpLongest = even
		}

		if len(tmpLongest) > len(longest) {
			longest = tmpLongest
		}
	}

	return longest
}

// O(n^2)
func getLongestPalindrome(str string, begin, end int) string {
	for begin >= 0 && end < len(str) {

		if str[begin] != str[end] {
			break
		}

		begin--
		end++
	}

	return str[begin+1 : end]
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
