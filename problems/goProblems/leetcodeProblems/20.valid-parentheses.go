package leetcodeProblems

/*
	@lc app=leetcode id=20 lang=golang

	[20] Valid Parentheses

	https://leetcode.com/problems/valid-parentheses/description/

	algorithms
	Easy (40.55%)
	Likes:    23609
	Dislikes: 1690
	Total Accepted:    4.5M
	Total Submissions: 11.1M
	Testcase Example:  '"()"'

	Given a string s containing just the characters '(', ')', '{', '}', '[' and
	']', determine if the input string is valid.

	An input string is valid if:


	Open brackets must be closed by the same type of brackets.
	Open brackets must be closed in the correct order.
	Every close bracket has a corresponding open bracket of the same type.



	Example 1:


	Input: s = "()"
	Output: true


	Example 2:


	Input: s = "()[]{}"
	Output: true


	Example 3:


	Input: s = "(]"
	Output: false
*/
/*
Accepted
98/98 cases passed (2 ms)
Your runtime beats 9.13 % of golang submissions
Your memory usage beats 43.85 % of golang submissions (2.3 MB)
*/

func (h *LCHandler) IsValidParens(s string) bool {
	return isValidParens(s)
}

// @lc code=start
// NOTE this problem has too generic a name, will need to rename
// to run uncomment below line and comment line 59

// func isValid(s string) bool {
func isValidParens(s string) bool {

	var stack Stack

	leftMap := map[rune]bool{'{': true, '[': true, '(': true}
	mapping := map[rune]rune{'}': '{', ']': '[', ')': '('}

	for _, char := range s {

		if _, foundLeft := leftMap[char]; foundLeft {
			stack.Push(char)
		}

		if expectedBracket, foundRight := mapping[char]; foundRight {

			if stack.IsEmpty() || expectedBracket != stack[len(stack)-1] {
				return false
			}
			stack.Pop()
		}
	}

	return stack.IsEmpty()
}

//type Stack []rune
// func (s *Stack) Push(val rune) {
// 	*s = append(*s, val)
// }

// func (s *Stack) Pop() (rune, bool) {
// 	if s.IsEmpty() {
// 		return 0, false
// 	}
// 	index := len(*s) - 1
// 	val := (*s)[index]
// 	*s = (*s)[:index]
// 	return val, true
// }

// func (s Stack) IsEmpty() bool {
// 	return len(s) == 0
// }

// @lc code=end
