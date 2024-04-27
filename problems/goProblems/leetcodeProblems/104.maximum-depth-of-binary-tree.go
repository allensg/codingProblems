package leetcodeProblems

/*
	@lc app=leetcode id=104 lang=golang

	[104] Maximum Depth of Binary Tree

	https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

	algorithms
	Easy (75.29%)
	Likes:    12616
	Dislikes: 223
	Total Accepted:    3M
	Total Submissions: 4M
	Testcase Example:  '[3,9,20,null,null,15,7]'

	Given the root of a binary tree, return its maximum depth.

	A binary tree's maximum depth is the number of nodes along the longest path
	from the root node down to the farthest leaf node.

	Example 1:
		Input: root = [3,9,20,null,null,15,7]
		Output: 3
	Example 2:
		Input: root = [1,null,2]
		Output: 2
*/
/*
Accepted
39/39 cases passed (0 ms)
Your runtime beats 100 % of golang submissions
Your memory usage beats 35.18 % of golang submissions (4.5 MB)
*/

func (h *LCHandler) MaxDepth(root *TreeNode) int {
	return maxDepth(root)
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// @lc code=end
