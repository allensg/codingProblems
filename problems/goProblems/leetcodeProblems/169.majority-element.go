package leetcodeProblems

/*
	@lc app=leetcode id=169 lang=golang

	[169] Majority Element

	https://leetcode.com/problems/majority-element/description/

	algorithms
	Easy (64.80%)
	Likes:    18724
	Dislikes: 592
	Total Accepted:    2.8M
	Total Submissions: 4.3M
	Testcase Example:  '[3,2,3]'

	Given an array nums of size n, return the majority element.

	The majority element is the element that appears more than ⌊n / 2⌋ times.
	You may assume that the majority element always exists in the array.

	Example 1:
		Input: nums = [3,2,3]
		Output: 3
	Example 2:
		Input: nums = [2,2,1,1,1,2,2]
		Output: 2
*/
/*
Accepted
51/51 cases passed (21 ms)
Your runtime beats 38.59 % of golang submissions
Your memory usage beats 50.17 % of golang submissions (6.7 MB)
*/
func (h *LCHandler) majorityElement(nums []int) int {
	return majorityElement(nums)
}

// @lc code=start
func majorityElement(nums []int) int {
	counts := make(map[int]int)
	majorityElementLen, currentHighestNumber := len(nums)/2, 0
	// this solution runs faster than 70% of other submitted answers. partly because
	// go is awesome and fast but also becaues of the majority element premise.
	// we don't acutally have to traverse the entire array to get the whole solution we just need to return
	// whatever number passes that number. so theoretically the cime complexity for this is O((n/2)+1) to O(n) depending on the ordering of the array.
	//  its possible we can't tell the majority until the very last element, or exactly 1 more than the middle element
	for _, value := range nums {
		currentCount := counts[value]
		if currentCount >= majorityElementLen {
			return value
		}
		newCount := currentCount + 1
		if newCount >= counts[currentHighestNumber] {
			currentHighestNumber = value
		}
		counts[value] = currentCount + 1
	}

	return currentHighestNumber
}

// @lc code=end
