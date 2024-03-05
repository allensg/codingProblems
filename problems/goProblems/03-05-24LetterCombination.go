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

	// iterate over the characters of the phone number
	for index, value := range phoneNumber {
		// pull the current number's letter value array
		currentDigitArray := digitMap[value]
		//loop over that array
		// ex: [a,b,c]
		for _, letter := range currentDigitArray {
			//loop over phonenumber[index+1:] value and add it to the combination
			for _, subDigit := range phoneNumber[index+1:] {
				subDigitArray := digitMap[int32(subDigit)]
				for _, subLetter := range subDigitArray {
					combinations = append(combinations, letter+subLetter)
				}
			}
		}
	}

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
