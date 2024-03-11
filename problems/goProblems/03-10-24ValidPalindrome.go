package goProblems

import (
	"regexp"
	"strings"
)

// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
// and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.
func (h *Handler) validPalindrome(s string) (string, bool) {
	// make lowercase
	s = strings.ToLower(s)
	// strip nonalpha
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = nonAlphanumericRegex.ReplaceAllString(s, "")

	for i := 0; len(s) > i; i++ {
		j := len(s) - i - 1
		if s[i] != s[j] {
			return "", false
		}
	}

	return "", true
}
