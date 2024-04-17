package leetcodeProblems

/*
	@lc app=leetcode id=189 lang=golang

	[189] Rotate Array

	https://leetcode.com/problems/rotate-array/description/

	algorithms
	Medium (40.52%)
	Likes:    17519
	Dislikes: 1923
	Total Accepted:    2.1M
	Total Submissions: 5.2M
	Testcase Example:  '[1,2,3,4,5,6,7]\n3'

	Given an integer array nums, rotate the array to the right by k steps, where
	k is non-negative.

	Example 1:
		Input: nums = [1,2,3,4,5,6,7], k = 3
		Output: [5,6,7,1,2,3,4]
		Explanation:
		rotate 1 steps to the right: [7,1,2,3,4,5,6]
		rotate 2 steps to the right: [6,7,1,2,3,4,5]
		rotate 3 steps to the right: [5,6,7,1,2,3,4]
	Example 2:
		Input: nums = [-1,-100,3,99], k = 2
		Output: [3,99,-1,-100]
		Explanation:
		rotate 1 steps to the right: [99,-1,-100,3]
		rotate 2 steps to the right: [3,99,-1,-100]
*/
/*
Accepted
38/38 cases passed (23 ms)
Your runtime beats 75.48 % of golang submissions
Your memory usage beats 11.5 % of golang submissions (8.6 MB)
*/
func (h *LCHandler) RotateArray(nums []int, k int) {
	rotate(nums, k)
}

// @lc code=start
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	toReturn := make([]int, 0, len(nums))
	// get k to end
	toReturn = append(toReturn, nums[len(nums)-k:]...)
	// get start to k
	toReturn = append(toReturn, nums[:len(nums)-k]...)
	// stick glue them together backwards
	// bing bang boom
	copy(nums, toReturn)
}

// @lc code=end
