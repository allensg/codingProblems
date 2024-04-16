package goProblems

// documenting medium problem from leetcode
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Given a string s, find the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(str string) int {
	repeating := make(map[rune]int) // map to store the index of characters
	longestLenSoFar, start := 0, 0

	for currentIndex, currentRune := range str {
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

// hash implementation. does not work exactly
// func lengthOfLongestSubstring(str string) int {
// 	repeating := make(map[rune]bool)

// 	currentSubString, longestLenSoFar := "", 0

// 	for _, currentRune := range str {
// 		runeStr := string(currentRune)
// 		used := repeating[currentRune]

// 		// if we find a repeating letter
// 		if used {
// 			fmt.Println("inside used")
// 			// check the length of this substring and store if it is the new longest.
// 			newLen := len(currentSubString)
// 			if newLen > longestLenSoFar {
// 				longestLenSoFar = newLen
// 			}
// 			// reset the list of repeating letters to create a new substring
// 			repeating = make(map[rune]bool)
// 			// set start of new current substring to current letter
// 			currentSubString = ""

// 		} else {
// 			fmt.Println("inside else")
// 			// if we haven't used the letter before
// 			// add it to the map
// 			repeating[currentRune] = true
// 			// and the substring
// 			currentSubString += runeStr
// 		}
// 	}

// 	return longestLenSoFar
// }
