package leetcodeProblems

/*
 @lc app=leetcode id=1 lang=golang

 [1] Two Sum

 https://leetcode.com/problems/two-sum/description/

 algorithms
 Easy (52.36%)
 Likes:    55835
 Dislikes: 1926
 Total Accepted:    12.9M
 Total Submissions: 24.6M
 Testcase Example:  '[2,7,11,15]\n9'

 Given an array of integers nums and an integer target, return indices of the
 two numbers such that they add up to target.

 You may assume that each input would have exactly one solution, and you may
 not use the same element twice.

 You can return the answer in any order.

 Example 1:
 	Input: nums = [2,7,11,15], target = 9
 	Output: [0,1]
 	Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
 Example 2:
 	Input: nums = [3,2,4], target = 6
 	Output: [1,2]
 Example 3:
 	Input: nums = [3,3], target = 6
 	Output: [0,1]
*/

func (h *LCHandler) TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

// @lc code=start
func twoSum(nums []int, target int) []int {
	return []int{}
}

// @lc code=end
