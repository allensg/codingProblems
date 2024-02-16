package goProblems

import "fmt"

// documenting medium problem from leetcode
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Given a string s, find the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(str string) int {
	repeating := make(map[rune]bool)

	currentSubString, longestLenSoFar := "", 0

	for _, currentRune := range str {
		runeStr := string(currentRune)
		used := repeating[currentRune]

		// if we find a repeating letter
		if used {
			fmt.Println("inside used")
			// check the length of this substring and store if it is the new longest.
			newLen := len(currentSubString)
			if newLen > longestLenSoFar {
				longestLenSoFar = newLen
			}
			// reset the list of repeating letters to create a new substring
			repeating = make(map[rune]bool)
			// set start of new current substring to current letter
			currentSubString = ""

		} else {
			fmt.Println("inside else")
			// if we haven't used the letter before
			// add it to the map
			repeating[currentRune] = true
			// and the substring
			currentSubString += runeStr
		}
	}

	return longestLenSoFar
}
