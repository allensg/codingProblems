package goProblems

import "sort"

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word
// or phrase, typically using all the original letters exactly once.
//EDIT: after solving the group anagrams problem its come to my attention that this is apparently, a stupid way to do things.
// a sort is faster

func (h *Handler) ValidAnagram(str1, str2 string) (returnString string, isAnagram bool) {

	var stringSort func(str string) string
	stringSort = func(str string) string {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		return string(runes)
	}

	sorted1, sorted2 := stringSort(str1), stringSort(str2)

	return "", sorted1 == sorted2
}

// old solution
// func (h *Handler) ValidAnagram(str1, str2 string) (returnString string, isAnagram bool) {

// 	var populateMap func(str string) map[rune]int
// 	populateMap = func(str string) map[rune]int {
// 		letterMap := map[rune]int{}
// 		for _, letter := range str {
// 			val, found := letterMap[letter]
// 			if found {
// 				letterMap[letter] = val + 1
// 			} else {
// 				letterMap[letter] = 0
// 			}
// 		}
// 		return letterMap
// 	}

// 	// traverse the first string, populate the map with the letters and their counts
// 	letterMap1 := populateMap(str1)
// 	// traverse the second string doing the same thing.
// 	letterMap2 := populateMap(str2)

// 	// the maps should be equal so if they're not the same length return
// 	if len(letterMap1) != len(letterMap2) {
// 		return "", false
// 	}

// 	// if they are equal length pick one, traverse it and compare each item to the item in the other list
// 	for key, map1Val := range letterMap1 {
// 		map2Val, map2Found := letterMap2[key]
// 		// if the current letter isn't found, or shows up an incorrect number of times it not an anagram, fail it
// 		if !map2Found || map2Val != map1Val {
// 			return "", false
// 		}
// 	}

// 	return "", true
// }
