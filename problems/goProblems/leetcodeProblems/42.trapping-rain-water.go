package leetcodeProblems

/*
	@lc app=leetcode id=42 lang=golang

	[42] Trapping Rain Water

	https://leetcode.com/problems/trapping-rain-water/description/

	algorithms
	Hard (61.87%)
	Likes:    31468
	Dislikes: 490
	Total Accepted:    2.1M
	Total Submissions: 3.3M
	Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'

	Given n non-negative integers representing an elevation map where the width
	of each bar is 1, compute how much water it can trap after raining.

 	Example 1:
		Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
		//
		//	            |
		//	    |       | |   |
		//	|   | |   | | | | | |
		//
		Output: 6
		Explanation: The above elevation map (black section) is represented by array
		[0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue
		section) are being trapped.
	Example 2:
		Input: height = [4,2,0,3,2,5]
		Output: 9
*/

// TODO
func (h *LCHandler) TrapRainwater(height []int) int {
	return trap(height)
}

// @lc code=start
func trap(height []int) int {
	return 0
}

// @lc code=end
