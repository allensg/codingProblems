package leetcodeProblems

import (
	"sort"
)

/*
	@lc app=leetcode id=49 lang=golang

	[49] Group Anagrams

	https://leetcode.com/problems/group-anagrams/description/

	algorithms
	Medium (68.45%)
	Likes:    18856
	Dislikes: 600
	Total Accepted:    2.8M
	Total Submissions: 4.1M
	Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'

	Given an array of strings strs, group the anagrams together. You can return
	the answer in any order.

	An Anagram is a word or phrase formed by rearranging the letters of a
	different word or phrase, typically using all the original letters exactly
	once.

	Example 1:
		Input: strs = ["eat","tea","tan","ate","nat","bat"]
		Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
	Example 2:
		Input: strs = [""]
		Output: [[""]]
	Example 3:
		Input: strs = ["a"]
		Output: [["a"]]
*/

/*
Accepted
126/126 cases passed (20 ms)
Your runtime beats 68.58 % of golang submissions
Your memory usage beats 51.45 % of golang submissions (7.9 MB)
*/

func (h *LCHandler) GroupAnagrams(strs []string) [][]string {
	return groupAnagrams(strs)
}

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	matched := map[string][]string{}
	// we can use the valid anagram from the last problem here.
	// iterate over the array
	for _, currentString := range strs {
		// take the current string and sort.
		// as we're grouping, if it matches a key it should match all the strings in the value array too
		sorted := sortString(currentString)
		subArr, found := matched[sorted]
		if found {
			// append value to the array
			matched[sorted] = append(subArr, currentString)
		} else {
			// if no match is found add a new item to the map with newly initialized subArray
			matched[sorted] = []string{currentString}
		}

	}

	toReturn := [][]string{}
	// build out return array
	for _, subArray := range matched {
		toReturn = append(toReturn, subArray)
	}
	return toReturn
}

func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// @lc code=end
