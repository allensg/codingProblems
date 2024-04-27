package leetcodeProblems

/*
	@lc app=leetcode id=13 lang=golang

	[13] Roman to Integer

	https://leetcode.com/problems/roman-to-integer/description/

	algorithms
	Easy (61.09%)
	Likes:    13794
	Dislikes: 899
	Total Accepted:    3.6M
	Total Submissions: 5.8M
	Testcase Example:  '"III"'

	Roman numerals are represented by seven different symbols: I, V, X, L, C, D
	and M.


	Symbol       Value
	I             1
	V             5
	X             10
	L             50
	C             100
	D             500
	M             1000

	For example, 2 is written as II in Roman numeral, just two ones added
	together. 12 is written as XII, which is simply X + II. The number 27 is
	written as XXVII, which is XX + V + II.

	Roman numerals are usually written largest to smallest from left to right.
	However, the numeral for four is not IIII. Instead, the number four is
	written as IV. Because the one is before the five we subtract it making
	four. The same principle applies to the number nine, which is written as IX.
	There are six instances where subtraction is used:

	I can be placed before V (5) and X (10) to make 4 and 9.
	X can be placed before L (50) and C (100) to make 40 and 90.
	C can be placed before D (500) and M (1000) to make 400 and 900.

	Given a roman numeral, convert it to an integer.

	Example 1:
		Input: s = "III"
		Output: 3
		Explanation: III = 3.
	Example 2:
		Input: s = "LVIII"
		Output: 58
		Explanation: L = 50, V= 5, III = 3.
	Example 3:
		Input: s = "MCMXCIV"
		Output: 1994
		Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/
/*
Accepted
3999/3999 cases passed (28 ms)
Your runtime beats 5.77 % of golang submissions
Your memory usage beats 8.24 % of golang submissions (4.1 MB)
*/
func (h *LCHandler) RomanToInt(s string) int {
	return romanToInt(s)
}

// @lc code=start
func romanToInt(s string) int {

	valueMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := valueMap[rune(s[0])]

	for index := 1; index < len(s); index++ {
		currentRune := rune(s[index])
		currentInt := valueMap[currentRune]

		// check to see if you need to add or subtract
		previousRune := rune(s[index-1])
		subtract := needsSubtraction(currentRune, previousRune)

		// if so add current (IV) and the previous (I), then subtract the previous (I)
		// (as we will have already added previous to the sum in the last iteration)
		if subtract {
			previousInt := valueMap[previousRune]
			toAdd := currentInt - previousInt
			sum = (sum - previousInt) + toAdd
		} else {
			sum = sum + currentInt
		}

	}

	return sum
}

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// takes in the current numeral and the previous one to determine if we need to subtract
func needsSubtraction(current, previous rune) bool {
	if (current == 'V' || current == 'X') && previous == 'I' {
		return true
	}

	if (current == 'L' || current == 'C') && previous == 'X' {
		return true
	}

	if (current == 'D' || current == 'M') && previous == 'C' {
		return true
	}

	return false
}

// @lc code=end
