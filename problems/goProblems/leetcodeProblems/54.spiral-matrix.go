package leetcodeProblems

/*
	@lc app=leetcode id=54 lang=golang

	[54] Spiral Matrix

	https://leetcode.com/problems/spiral-matrix/description/

	algorithms
	Medium (49.33%)
	Likes:    14422
	Dislikes: 1272
	Total Accepted:    1.4M
	Total Submissions: 2.8M
	Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'

	Given an m x n matrix, return all elements of the matrix in spiral order.

	Example 1:
	Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
	Output: [1,2,3,6,9,8,7,4,5]
	Example 2:
	Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
	Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/
/*
Accepted
25/25 cases passed (1 ms)
Your runtime beats 79.66 % of golang submissions
Your memory usage beats 6.1 % of golang submissions (2.2 MB)
*/

func (h *LCHandler) SpiralOrder(matrix [][]int) []int {
	return spiralOrder(matrix)
}

// @lc code=start
func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return []int{}
	}

	var result []int
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// Traverse top row
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse rightmost column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Ensure we haven't crossed boundaries
		if top <= bottom {
			// Traverse bottom row
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			// Traverse leftmost column
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

// @lc code=end
