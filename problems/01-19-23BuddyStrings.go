package problems

import "fmt"

// Given two strings A and B of lowercase letters, return true if and
// only if we can swap two letters in A so that the result equals B.

// input: A="ab", B="ba"
// output: true
func (h *Handler) BuddyStrings(input []string) (returnString string, answer bool) {
	a, b := []rune(input[0]), []rune(input[1])
	returnString = fmt.Sprintf("Strings %s and %s are compatible", a, b)
	if len(a) != len(b) {
		return fmt.Sprintf("Strings %s and %s are not of equal length", a, b), false
	}

	for i, k := range a {

	}
	// if a == b && len(a) > len(set(a)) {
	// 	return returnString, true
	// }

	return returnString, answer
}
