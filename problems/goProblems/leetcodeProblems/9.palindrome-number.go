package leetcodeProblems

/*
	@lc app=leetcode id=9 lang=golang

	[9] Palindrome Number

	https://leetcode.com/problems/palindrome-number/description/

	algorithms
	Easy (56.19%)
	Likes:    12293
	Dislikes: 2701
	Total Accepted:    4.5M
	Total Submissions: 8M
	Testcase Example:  '121'

	Given an integer x, return true if x is a palindrome, and false
	otherwise.

	Example 1:
		Input: x = 121
		Output: true
		Explanation: 121 reads as 121 from left to right and from right to left.
	Example 2:
		Input: x = -121
		Output: false
		Explanation: From left to right, it reads -121. From right to left, it
		becomes 121-. Therefore it is not a palindrome.
	Example 3:
		Input: x = 10
		Output: false
		Explanation: Reads 01 from right to left. Therefore it is not a
		palindrome.
*/
/*
Accepted
11511/11511 cases passed (14 ms)
Your runtime beats 45.06 % of golang submissions
Your memory usage beats 97.28 % of golang submissions (4.3 MB)
*/
func (h *LCHandler) IsNumberPalindrome(x int) bool {
	return isNumberPalindrome(x)
}

// @lc code=start
// NOTE this is a followup to another problem, will need to rename
// to run uncomment below line and comment line 44
// func isPalindrome(x int) bool {
func isNumberPalindrome(x int) bool {
	// negatives ain't palindromes
	if x < 0 {
		return false
	}

	num, reversed := x, 0
	for num != 0 {
		// reverse the number
		// shift the digit by * 10
		// pop the next digit by % 10
		// join by addition
		reversed = (10 * reversed) + (num % 10)
		// shift num down by division until we run out of digits (num == 0)
		num = num / 10
	}
	// if the number equals its reverse, its a palindrome
	toReturn := (x == reversed)

	return toReturn
}

// @lc code=end
