package leetcodeProblems

/*
	@lc app=leetcode id=6 lang=golang

	[6] Zigzag Conversion

	https://leetcode.com/problems/zigzag-conversion/description/

	algorithms
	Medium (47.80%)
	Likes:    7435
	Dislikes: 14364
	Total Accepted:    1.3M
	Total Submissions: 2.7M
	Testcase Example:  '"PAYPALISHIRING"\n3'

	The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
	of rows like this: (you may want to display this pattern in a fixed font for
	better legibility)

	P   A   H   N
	A P L S I I G
	Y   I   R

	And then read line by line: "PAHNAPLSIIGYIR"
	Write the code that will take a string and make this conversion given a
	number of rows:
	string convert(string s, int numRows);

	Example 1:
		Input: s = "PAYPALISHIRING", numRows = 3
		Output: "PAHNAPLSIIGYIR"
	Example 2:
		Input: s = "PAYPALISHIRING", numRows = 4
		Output: "PINALSIGYAHRPI"
		Explanation:
		P     I    N
		A   L S  I G
		Y A   H R
		P     I
	Example 3:
		Input: s = "A", numRows = 1
		Output: "A"

*/

func (h *LCHandler) ZigZagConversion(s string, numRows int) string {
	return convert(s, numRows)
}

// @lc code=start
// TODO
func convert(s string, numRows int) string {
	ZigZag := StringToZigZag(s, numRows)
	scrambledString := ZigZagToString(ZigZag)

	return scrambledString
}

// takes in a string and stores it in columns like so
// Input: s = "PAYPALISHIRING", numRows = 4
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// where index column (k) 0 is fully populated
// then each column after that is staggered such that column is placed at row - 1 until k%len(str) is satisfied
func StringToZigZag(str string, lines int) *[][]rune {
	strLen := len(str)

	// build array
	ZigZag := make([][]rune, lines)
	for index, _ := range ZigZag {
		ZigZag[index] = make([]rune, strLen)
	}

	// in this case the index is the current row of the column we are on.
	// the mod is the column we are on
	// the next value should be the current length of the column - the current remainder.
	// ie 4%(0,4) = 0, column (0,4) should hit the if statement
	// 4%(1,5) = 1 column (1,5) should have all spaces with the next letter at row columns-1
	// 4%(2,6) = 2 column (2,6) should have all spaces with the next letter at row columns-2
	// 4%(3,7) = 3 column (3,7) should have all spaces with the next letter at row columns-3
	for index, currentRune := range str {
		// on mod zero letters store all letters in a column
		mod := strLen % index
		if mod == 0 {
			ZigZag[mod][index] = currentRune
		} else {
			ZigZag[mod][strLen-index] = currentRune
		}
	}

	return &ZigZag
}

func ZigZagToString(*[][]rune) string {
	return ""
}

// @lc code=end
