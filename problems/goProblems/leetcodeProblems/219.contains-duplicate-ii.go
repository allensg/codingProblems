package leetcodeProblems

/*
	@lc app=leetcode id=219 lang=golang

	[219] Contains Duplicate II

	https://leetcode.com/problems/contains-duplicate-ii/description/

	algorithms
	Easy (45.00%)
	Likes:    6012
	Dislikes: 3055
	Total Accepted:    918K
	Total Submissions: 2M
	Testcase Example:  '[1,2,3,1]\n3'

	Given an integer array nums and an integer k, return true if there are two
	distinct indices i and j in the array such that nums[i] == nums[j] and abs(i
	- j) <= k.

	Example 1:
		Input: nums = [1,2,3,1], k = 3
		Output: true
	Example 2:
		Input: nums = [1,0,1,1], k = 1
		Output: true
	Example 3:
		Input: nums = [1,2,3,1,2,3], k = 2
		Output: false
*/
/*
Accepted
58/58 cases passed (73 ms)
Your runtime beats 97.81 % of golang submissions
Your memory usage beats 53.83 % of golang submissions (12.4 MB)
*/

func (h *LCHandler) ContainsNearbyDuplicate(nums []int, k int) bool {
	return containsNearbyDuplicate(nums, k)
}

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	// map values to their indexes
	indexMap := map[int]int{}

	for currentIndex, value := range nums {
		// pull index map
		mapIndex, found := indexMap[value]
		// if we have a previous instance and
		// that instance's index is within the range of the current index
		if found && currentIndex-mapIndex <= k {
			return true
		}
		// update map
		indexMap[value] = currentIndex
	}

	return false
}

// @lc code=end
