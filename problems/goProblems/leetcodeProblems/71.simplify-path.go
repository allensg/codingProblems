package leetcodeProblems

import (
	"strings"
)

/*
	@lc app=leetcode id=71 lang=golang

	[71] Simplify Path

	https://leetcode.com/problems/simplify-path/description/

	algorithms
	Medium (42.27%)
	Likes:    5427
	Dislikes: 1254
	Total Accepted:    736K
	Total Submissions: 1.7M
	Testcase Example:  '"/home/"'

	Given a string path, which is an absolute path (starting with a slash '/')
	to a file or directory in a Unix-style file system, convert it to the
	simplified canonical path.

	In a Unix-style file system, a period '.' refers to the current directory, a
	double period '..' refers to the directory up a level, and any multiple
	consecutive slashes (i.e. '//') are treated as a single slash '/'. For this
	problem, any other format of periods such as '...' are treated as
	file/directory names.

	The canonical path should have the following format:

	The path starts with a single slash '/'.
	Any two directories are separated by a single slash '/'.
	The path does not end with a trailing '/'.
	The path only contains the directories on the path from the root directory
	to the target file or directory (i.e., no period '.' or double period
	'..')

	Return the simplified canonical path.

	Example 1:
		Input: path = "/home/"
		Output: "/home"
		Explanation: Note that there is no trailing slash after the last directory
		name.
	Example 2:
		Input: path = "/../"
		Output: "/"
		Explanation: Going one level up from the root directory is a no-op, as the
		root level is the highest level you can go.
	Example 3:
		Input: path = "/home//foo/"
		Output: "/home/foo"
		Explanation: In the canonical path, multiple consecutive slashes are
		replaced by a single one.
*/
/*
Accepted
259/259 cases passed (3 ms)
Your runtime beats 54.66 % of golang submissions
Your memory usage beats 59 % of golang submissions (3.4 MB)
*/

func (h *LCHandler) SimplifyPath(path string) string {
	return simplifyPath(path)
}

// @lc code=start
func simplifyPath(path string) string {
	var stack StringStack

	split := strings.Split(path, "/")

	for _, val := range split {

		switch val {
		// every time we hit a / we add that directory to the stack
		case "", ".":
			continue
		// everytime we hit a .. we pop a directory off the stack
		case "..":
			if !stack.IsEmpty() {
				stack.Pop()
			}
		default:
			stack.Push(val)
		}

	}
	// actually because the stack is just an array we can just join strings.join /
	toReturn := "/" + strings.Join(stack, "/")
	return toReturn
}

// type StringStack []string

// // string
// func (s *StringStack) Push(val string) {
// 	*s = append(*s, val)
// }

// func (s *StringStack) Pop() (string, bool) {
// 	if s.IsEmpty() {
// 		return "", false
// 	}
// 	index := len(*s) - 1
// 	val := (*s)[index]
// 	*s = (*s)[:index]
// 	return val, true
// }

// func (s StringStack) IsEmpty() bool {
// 	return len(s) == 0
// }

// @lc code=end
