package leetcodeProblems

/*
 @lc app=leetcode id=3 lang=golang

 [3] Longest Substring Without Repeating Characters

 https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

 algorithms
 Medium (34.72%)
 Likes:    39099
 Dislikes: 1835
 Total Accepted:    5.7M
 Total Submissions: 16.3M
 Testcase Example:  '"abcabcbb"'

 Given a string s, find the length of the longest substring without repeating
 characters.

 Example 1:
	Input: s = "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3.
 Example 2:
	Input: s = "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.
Example 3:
	Input: s = "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.
	Notice that the answer must be a substring, "pwke" is a subsequence and not
	a substring.
*/

/*
  Accepted
	987/987 cases passed (4 ms)
	Your runtime beats 83.32 % of golang submissions
	Your memory usage beats 44.26 % of golang submissions (3.4 MB)
*/

func (h *LCHandler) LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	repeating := make(map[rune]int) // map to store the index of characters
	longestLenSoFar, start := 0, 0

	for currentIndex, currentRune := range s {
		runeIndex, found := repeating[currentRune]
		if found && runeIndex >= start {
			// If the character is repeating within the current substring
			// Update the start index to the next index of the repeating character
			start = runeIndex + 1
		}
		// Update the index of the current character
		repeating[currentRune] = currentIndex
		// Calculate the length of the current substring
		currentLen := currentIndex - start + 1
		if currentLen > longestLenSoFar {
			longestLenSoFar = currentLen
		}
	}
	return longestLenSoFar
}

// @lc code=end
