package leetcodeProblems

/*
	@lc app=leetcode id=637 lang=golang

	[637] Average of Levels in Binary Tree

	https://leetcode.com/problems/average-of-levels-in-binary-tree/description/

	algorithms
	Easy (72.45%)
	Likes:    5236
	Dislikes: 325
	Total Accepted:    506.4K
	Total Submissions: 697.8K
	Testcase Example:  '[3,9,20,null,null,15,7]'

	Given the root of a binary tree, return the average value of the nodes on
	each level in the form of an array. Answers within 10^-5 of the actual
	answer will be accepted.

	Example 1:
		Input: root = [3,9,20,null,null,15,7]
		Output: [3.00000,14.50000,11.00000]
		Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5,
		and on level 2 is 11.
		Hence return [3, 14.5, 11].
	Example 2:
		Input: root = [3,9,20,15,7]
		Output: [3.00000,14.50000,11.00000]
*/

func (h *LCHandler) AverageOfLevels(root *TreeNode) []float64 {
	return averageOfLevels(root)
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
func averageOfLevels(root *TreeNode) []float64 {
	toReturn := []float64{}

	if root == nil {
		return toReturn
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sum := 0.0
		levelLen := len(queue)

		for _, current := range queue {
			queue = queue[1:]
			sum = sum + float64(current.Val)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		average := sum / float64(levelLen)
		toReturn = append(toReturn, average)

	}

	return toReturn
}

// @lc code=end
