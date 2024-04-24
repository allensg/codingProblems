package leetcodeProblems

/*
 	@lc app=leetcode id=392 lang=golang

	[392] Is Subsequence

	https://leetcode.com/problems/is-subsequence/description/

	algorithms
	Easy (47.93%)
	Likes:    9427
	Dislikes: 514
	Total Accepted:    1.4M
	Total Submissions: 2.9M
	Testcase Example:  '"abc"\n"ahbgdc"'

	Given two strings s and t, return true if s is a subsequence of t, or false
	otherwise.

	A subsequence of a string is a new string that is formed from the original
	string by deleting some (can be none) of the characters without disturbing
	the relative positions of the remaining characters. (i.e., "ace" is a
	subsequence of "abcde" while "aec" is not).

	Example 1:
		Input: s = "abc", t = "ahbgdc"
		Output: true
	Example 2:
		Input: s = "axc", t = "ahbgdc"
		Output: false
*/
/*
Accepted
20/20 cases passed (1 ms)
Your runtime beats 84.13 % of golang submissions
Your memory usage beats 36.58 % of golang submissions (2.3 MB)
*/

func (h *LCHandler) isSubsequence(s string, t string) bool {
	return isSubsequence(s, t)
}

// @lc code=start
func isSubsequence(s string, t string) bool {

	lastFoundAt := 0
	for _, sRune := range s {
		foundSubMatch := false
		for tIndex := lastFoundAt; tIndex < len(t); tIndex++ {
			if rune(t[tIndex]) == sRune {
				lastFoundAt = tIndex + 1
				foundSubMatch = true
				break
			}
		}

		if !foundSubMatch {
			return false
		}
	}

	return true
}

// @lc code=end
