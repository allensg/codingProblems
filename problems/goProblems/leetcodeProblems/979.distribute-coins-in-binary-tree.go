package leetcodeProblems

/*
	@lc app=leetcode id=979 lang=golang

	[979] Distribute Coins in Binary Tree

	https://leetcode.com/problems/distribute-coins-in-binary-tree/description/

	algorithms
	Medium (72.57%)
	Likes:    5600
	Dislikes: 222
	Total Accepted:    180.4K
	Total Submissions: 235.2K
	Testcase Example:  '[3,0,0]'

	You are given the root of a binary tree with n nodes where each node in the
	tree has node.val coins. There are n coins in total throughout the whole
	tree.

	In one move, we may choose two adjacent nodes and move one coin from one
	node to another. A move may be from parent to child, or from child to
	parent.

	Return the minimum number of moves required to make every node have exactly
	one coin.

	Example 1:
		Input: root = [3,0,0]
		Output: 2
		Explanation: From the root of the tree, we move one coin to its left child,
		and one coin to its right child.
	Example 2:
		Input: root = [0,3,0]
		Output: 3
		Explanation: From the left child of the root, we move two coins to the root
		[taking two moves]. Then, we move one coin from the root of the tree to the
		right child.
*/

// @lc code=start
func (h *LCHandler) DistributeCoins(root *TreeNode) int {
	return distributeCoins(root)
}

func distributeCoins(root *TreeNode) int {
	moves := 0
	moves = moves + processNodes(root.Left, root)
	moves = moves + processNodes(root.Right, root)
	return moves
}

// @lc code=end
