package leetcode

import (
	"strings"
	"unicode"
)

// First Brute force method
func isPalindrome(s string) bool {
	cleanedString := cleanString(s)
	middle := len(cleanedString) / 2
	endPointer := len(cleanedString) - 1

	for i := 0; i < middle; i++ {

		if cleanedString[i] != cleanedString[endPointer] {
			return false
		}
		endPointer -= 1
	}
	return true
}

func cleanString(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, strings.TrimSpace(str))
}
