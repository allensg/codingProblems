package leetcodeProblems

import (
	"regexp"
	"strings"
)

/*
	@lc app=leetcode id=125 lang=golang

	[125] Valid Palindrome

 	https://leetcode.com/problems/valid-palindrome/description/

 	algorithms
 	Easy (47.34%)
 	Likes:    9076
 	Dislikes: 8285
 	Total Accepted:    2.9M
 	Total Submissions: 6.1M
 	Testcase Example:  '"A man, a plan, a canal: Panama"'

 	A phrase is a palindrome if, after converting all uppercase letters into
 	lowercase letters and removing all non-alphanumeric characters, it reads the
 	same forward and backward. Alphanumeric characters include letters and
 	numbers.

 	Given a string s, return true if it is a palindrome, or false otherwise.

 	Example 1:
		Input: s = "A man, a plan, a canal: Panama"
		Output: true
		Explanation: "amanaplanacanalpanama" is a palindrome.
 	Example 2:
		Input: s = "race a car"
		Output: false
		Explanation: "raceacar" is not a palindrome.
 	Example 3:
		Input: s = " "
		Output: true
		Explanation: s is an empty string "" after removing non-alphanumeric
		characters.
		Since an empty string reads the same forward and backward, it is a
		palindrome.
*/
/*
	Accepted
		485/485 cases passed (7 ms)
		Your runtime beats 31.96 % of golang submissions
		Your memory usage beats 34.15 % of golang submissions (4.3 MB)
*/
func (h *LCHandler) IsPalindrome(s string) bool {
	return isPalindrome(s)
}

// @lc code=start
func isPalindrome(s string) bool {
	// make lowercase
	s = strings.ToLower(s)
	// strip nonalpha
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = nonAlphanumericRegex.ReplaceAllString(s, "")
	// traverse the array, starting at the ends and moving towards the middle
	// if anything doesn't line up, return
	for i := 0; len(s) > i; i++ {
		j := len(s) - i - 1
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

// @lc code=end
