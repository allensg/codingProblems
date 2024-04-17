package leetcodeProblems

/*
	@lc app=leetcode id=17 lang=golang

	[17] Letter Combinations of a Phone Number

	https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/

	algorithms
	Medium (60.30%)
	Likes:    18174
	Dislikes: 974
	Total Accepted:    2M
	Total Submissions: 3.3M
	Testcase Example:  '"23"'

	Given a string containing digits from 2-9 inclusive, return all possible
	letter combinations that the number could represent. Return the answer in
	any order.

	A mapping of digits to letters (just like on the telephone buttons) is given
	below. Note that 1 does not map to any letters.

	Example 1:
		Input: digits = "23"
	 	Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	Example 2:
		Input: digits = ""
	 	Output: []
	Example 3:
		Input: digits = "2"
	 	Output: ["a","b","c"]
*/

/*
Accepted
25/25 cases passed (1 ms)
Your runtime beats 79.23 % of golang submissions
Your memory usage beats 16.54 % of golang submissions (2.4 MB)
*/
func (h *LCHandler) LetterCombinations(digits string) []string {
	return letterCombinations(digits)
}

// @lc code=start
func letterCombinations(digits string) []string {
	digitMap := makeDigitMap()
	combinations := []string{}

	if len(digits) == 0 {
		return combinations
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

	traverse("", digits)
	return combinations
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

// @lc code=end
