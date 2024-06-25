package leetcodeProblems

/*
	@lc app=leetcode id=1038 lang=golang

	[1038] Binary Search Tree to Greater Sum Tree

	https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/description/

	algorithms
	Medium (88.08%)
	Likes:    4224
	Dislikes: 160
	Total Accepted:    269.4K
	Total Submissions: 305.9K
	Testcase Example:  '[4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]'

	Given the root of a Binary Search Tree (BST), convert it to a Greater Tree
	such that every key of the original BST is changed to the original key plus
	the sum of all keys greater than the original key in BST.

	As a reminder, a binary search tree is a tree that satisfies these
	constraints:

	The left subtree of a node contains only nodes with keys less than the
	node's key.
	The right subtree of a node contains only nodes with keys greater than the
	node's key.
	Both the left and right subtrees must also be binary search trees.

	Example 1:
		Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
		Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
	Example 2:
		Input: root = [0,null,1]
		Output: [1,null,1]
*/

/*
Accepted
30/30 cases passed (1 ms)
Your runtime beats 67.16 % of golang submissions
Your memory usage beats 17.91 % of golang submissions (2.4 MB)
*/
func (lh *LCHandler) BstToGst(root *TreeNode) *TreeNode {
	return bstToGst(root)
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		// go all the way right
		traverse(root.Right)
		// keep track of the sum so far and update
		sum = sum + root.Val
		root.Val = sum
		// apply sum to the left and continue.
		traverse(root.Left)
	}

	traverse(root)
	return root
}

// @lc code=end
