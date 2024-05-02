package leetcodeProblems

/*
	@lc app=leetcode id=28 lang=golang

	[28] Find the Index of the First Occurrence in a String

	https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/

	algorithms
	Easy (42.20%)
	Likes:    5616
	Dislikes: 383
	Total Accepted:    2.4M
	Total Submissions: 5.7M
	Testcase Example:  '"sadbutsad"\n"sad"'

	Given two strings needle and haystack, return the index of the first
	occurrence of needle in haystack, or -1 if needle is not part of
	haystack.

	Example 1:
		Input: haystack = "sadbutsad", needle = "sad"
		Output: 0
		Explanation: "sad" occurs at index 0 and 6.
		The first occurrence is at index 0, so we return 0.
	 Example 2:
		Input: haystack = "leetcode", needle = "leeto"
		Output: -1
		Explanation: "leeto" did not occur in "leetcode", so we return -1.
*/

func (h *LCHandler) StrStr(haystack string, needle string) int {
	return strStr(haystack, needle)
}

// @lc code=start
func strStr(haystack string, needle string) int {
	occurrenceIndex := 0
	needleIndex := 0
	for hIndex, hValue := range haystack {

		if hValue == rune(needle[needleIndex]) {
			needleIndex = needleIndex + 1
		}

	}

	return -1
}

// @lc code=end
