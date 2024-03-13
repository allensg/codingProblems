package goProblems

import "strings"

// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

// A mapping of digits to letters (just like on the telephone buttons) is given below.
// Note that 1 does not map to any letters.
// 2 - abc, 3-def, 4-ghi, 5-jkl, 6-mno, 7-pqrs, 8-tuv, 9-wxyz

// ex
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

// i really don't like that this has 4 nested loops. maybe something with recursion. wouldn't necessarily
// be faster as we do need every combination but it would be visually cleaner
func (h *Handler) LetterCombinations(phoneNumber string) (returnString string) {
	digitMap := makeDigitMap()
	combinations := []string{}

	if len(phoneNumber) == 0 {
		return ""
	}
	// 	Define Recursive Function: Utilize a recursive function backtrack that takes the current combination and the next digits to explore.
	var traverse func(current string, next string)
	traverse = func(current string, next string) {
		// Termination Condition: If there are no more digits to explore, append the current combination to the result.
		// Exploration: Iterate through the corresponding letters of the next digit and recursively explore the remaining digits.
		if next == "" {
			combinations = append(combinations, current)
		} else {
			currentDigitArray := digitMap[rune(next[0])]
			for _, letter := range currentDigitArray {
				traverse(current+string(letter), next[1:])
			}
		}
	}

	traverse("", phoneNumber)
	return strings.Join(combinations, " ")
}

func makeDigitMap() map[rune][]string {
	digitMap := make(map[rune][]string)
	digitMap['2'] = []string{"a", "b", "c"}
	digitMap['3'] = []string{"d", "e", "f"}
	digitMap['4'] = []string{"g", "h", "i"}
	digitMap['5'] = []string{"j", "k", "l"}
	digitMap['6'] = []string{"m", "n", "o"}
	digitMap['7'] = []string{"p", "q", "r", "s"}
	digitMap['8'] = []string{"t", "u", "v"}
	digitMap['9'] = []string{"w", "x", "y", "z"}

	return digitMap
}
