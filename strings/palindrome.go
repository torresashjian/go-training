package strings

import ()

func IsPalindrome(str string) bool {
	isPalindrome := true
	for i := 0; i < len(str); i++ {
		forwardLetter := str[i]
		backwardLetter := str[len(str)-1-i]
		if forwardLetter != backwardLetter {
			isPalindrome = false
			break
		}
	}
	return isPalindrome
}
