package leetcodeProblems

/*
	@lc app=leetcode id=55 lang=golang

	[55] Jump Game

	https://leetcode.com/problems/jump-game/description/

	algorithms
	Medium (38.58%)
	Likes:    19237
	Dislikes: 1232
	Total Accepted:    2M
	Total Submissions: 5.1M
	Testcase Example:  '[2,3,1,1,4]'

	You are given an integer array nums. You are initially positioned at the
	array's first index, and each element in the array represents your maximum
	jump length at that position.

	Return true if you can reach the last index, or false otherwise.

	Example 1:
		Input: nums = [2,3,1,1,4]
		Output: true
		Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last
		index.
	Example 2:
		Input: nums = [3,2,1,0,4]
		Output: false
		Explanation: You will always arrive at index 3 no matter what. Its maximum
		jump length is 0, which makes it impossible to reach the last index.
*/

func (lc *LCHandler) CanJump(nums []int) bool {
	return canJump(nums)
}

// @lc code=start
func canJump(nums []int) bool {
	// dfs jumping with a stack
	return true
}

// @lc code=end
