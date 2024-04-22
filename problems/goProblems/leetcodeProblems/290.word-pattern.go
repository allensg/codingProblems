package leetcodeProblems

import (
	"strings"
)

/*
	@lc app=leetcode id=290 lang=golang

	[290] Word Pattern

	https://leetcode.com/problems/word-pattern/description/

	algorithms
	Easy (42.01%)
	Likes:    7151
	Dislikes: 983
	Total Accepted:    697.2K
	Total Submissions: 1.7M
	Testcase Example:  '"abba"\n"dog cat cat dog"'

	Given a pattern and a string s, find if s follows the same pattern.

	Here follow means a full match, such that there is a bijection between a
	letter in pattern and a non-empty word in s.

	Example 1:
		Input: pattern = "abba", s = "dog cat cat dog"
		Output: true
	Example 2:
		Input: pattern = "abba", s = "dog cat cat fish"
		Output: false
	Example 3:
		Input: pattern = "aaaa", s = "dog cat cat dog"
		Output: false

	Constraints:
	s contains only lowercase English letters and spaces ' '.
	s does not contain any leading or trailing spaces.
	All the words in s are separated by a single space.
*/
/*
Accepted
42/42 cases passed (1 ms)
Your runtime beats 83.47 % of golang submissions
Your memory usage beats 35.02 % of golang submissions (2.3 MB)
*/

func (h *LCHandler) WordPattern(pattern string, s string) bool {
	return wordPattern(pattern, s)
}

// @lc code=start
func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(words) != len(pattern) {
		return false
	}
	patternToWord := make(map[rune]string)
	wordToPattern := make(map[string]rune)
	for index, char := range pattern {
		word := words[index]
		// check if our pattern letter already has a word mapped to it
		if p2w, ok := patternToWord[char]; ok {
			if p2w != word {
				return false
			}
		} else {
			patternToWord[char] = word
		}
		if w2p, ok := wordToPattern[word]; ok {
			if w2p != char {
				return false
			}
		} else {
			wordToPattern[word] = char
		}
	}
	return true
}

// @lc code=end
