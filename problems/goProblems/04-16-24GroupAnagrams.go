package goProblems

import "sort"

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:

// Input: strs = [""]
// Output: [[""]]
// Example 3:

// Input: strs = ["a"]
// Output: [["a"]]

func (h *Handler) GroupedAnagrams(strs []string) (returnString string, groupedAnagrams [][]string) {
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
	for _, subArray := range matched {
		toReturn = append(toReturn, subArray)
	}
	return "", toReturn
}

// Helper function to sort a string
func sortString(str string) string {
	runes := []rune(str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// original idea but too slow. need to switch to a hash implementation where toreturn is a map of base strings and arrays,
// then check the map to build the sub arrays and build the return array
// func (h *Handler) GroupedAnagrams(strs []string) (returnString string, groupedAnagrams [][]string) {
// 	toReturn := [][]string{}
// 	alreadyMatched := map[string]bool{}
// 	// we can use the valid palindrome from the last problem here.
// 	// iterate over the array
// 	for topIndex, currentString := range strs {
// 		// check to see if we have already matched this item so we don't end up with
// 		// duplicates in the final array
// 		if _, found := alreadyMatched[currentString]; found {
// 			continue
// 		}

// 		// establish new sub array for this entry.
// 		subArray := []string{currentString}
// 		// check all subsequent entries for anagrams
// 		for _, currentSubString := range strs[topIndex+1:] {

// 			// if match, append to sub array and mark that string as matched.
// 			if _, match := h.ValidAnagram(currentString, currentSubString); match {
// 				subArray = append(subArray, currentSubString)
// 				alreadyMatched[currentSubString] = true
// 			}
// 		}
// 		// append subArray to groupedAnagrams
// 		toReturn = append(toReturn, subArray)
// 	}
// 	return "", toReturn
// }
