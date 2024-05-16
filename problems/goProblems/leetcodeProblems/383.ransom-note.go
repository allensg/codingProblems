package leetcodeProblems

import "fmt"

/*
	@lc app=leetcode id=383 lang=golang

	[383] Ransom Note

	https://leetcode.com/problems/ransom-note/description/

	algorithms
	Easy (61.24%)
	Likes:    4910
	Dislikes: 496
	Total Accepted:    1.1M
	Total Submissions: 1.9M
	Testcase Example:  '"a"\n"b"'

	Given two strings ransomNote and magazine, return true if ransomNote can be
	constructed by using the letters from magazine and false otherwise.

	Each letter in magazine can only be used once in ransomNote.

	Example 1:
		Input: ransomNote = "a", magazine = "b"
		Output: false
	Example 2:
		Input: ransomNote = "aa", magazine = "ab"
		Output: false
	Example 3:
		Input: ransomNote = "aa", magazine = "aab"
		Output: true
*/
/*

 */

func (h *LCHandler) CanConstruct(ransomNote string, magazine string) bool {
	return canConstruct(ransomNote, magazine)
}

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {

	noteMap := generateMap(ransomNote)
	magMap := generateMap(magazine)

	for _, r := range ransomNote {
		noteCount, _ := noteMap[r]
		magCount, magFound := magMap[r]
		if !magFound || magCount < noteCount {
			return false
		}
	}

	return true
}

func generateMap(str string) map[rune]int {
	toReturn := map[rune]int{}

	for _, letter := range str {
		count, found := toReturn[letter]
		if found {
			toReturn[letter] = count + 1
		} else {
			toReturn[letter] = 1
		}

	}
	fmt.Printf("map: %v\n", toReturn)
	return toReturn
}

// @lc code=end
