package leetcodeProblems

/*
	@lc app=leetcode id=78 lang=golang

	[78] Subsets

	https://leetcode.com/problems/subsets/description/

	algorithms
	Medium (77.39%)
	Likes:    17024
	Dislikes: 276
	Total Accepted:    1.9M
	Total Submissions: 2.5M
	Testcase Example:  '[1,2,3]'

	Given an integer array nums of unique elements, return all possible subsets
	(the power set).

	The solution set must not contain duplicate subsets. Return the solution in
	any order.

	Example 1:
		Input: nums = [1,2,3]
		Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
	Example 2:
		Input: nums = [0]
		Output: [[],[0]]
*/

func (h *LCHandler) Subsets(nums []int) [][]int {
	return subsets(nums)
}

// @lc code=start
func subsets(nums []int) [][]int {

	toReturn := [][]int{[]int{}}

	for numI, num := range nums {
		toReturn = append(toReturn, []int{num})
		subArr := []int{num}
		for _, subNum := range nums[numI+1:] {
			// add pairs
			pair := []int{num, subNum}
			toReturn = append(toReturn, pair)

			subArr = append(subArr, subNum)
			toReturn = append(toReturn, subArr)
		}
		// toReturn = append(toReturn, subArr)
	}

	toReturn = append(toReturn, nums)
	return toReturn
}

// @lc code=end
