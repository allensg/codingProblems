package leetcodeProblems

import "sort"

/*
	@lc app=leetcode id=242 lang=golang

	[242] Valid Anagram

	https://leetcode.com/problems/valid-anagram/description/

	algorithms
	Easy (64.45%)

	Likes:    11880
	Dislikes: 390
	Total Accepted:    3.4M
	Total Submissions: 5.3M
	Testcase Example:  '"anagram"\n"nagaram"'

	Given two strings s and t, return true if t is an anagram of s, and false
	otherwise.

	An Anagram is a word or phrase formed by rearranging the letters of a
	different word or phrase, typically using all the original letters exactly
	once.

	Example 1:
		Input: s = "anagram", t = "nagaram"
		Output: true
	Example 2:
		Input: s = "rat", t = "car"
		Output: false

*/

func (h *LCHandler) IsAnagram(s string, t string) bool {
	return isAnagram(s, t)
}

// @lc code=start
func isAnagram(s string, t string) bool {
	sortString := func(str string) string {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		return string(runes)
	}
	return sortString(s) == sortString(t)
}

// @lc code=end
