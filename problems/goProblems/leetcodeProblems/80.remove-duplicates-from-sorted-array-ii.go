package leetcodeProblems

import "slices"

/*
	@lc app=leetcode id=80 lang=golang

	[80] Remove Duplicates from Sorted Array II

	https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/

	algorithms
	Medium (58.03%)
	Likes:    6659
	Dislikes: 1275
	Total Accepted:    914.4K
	Total Submissions: 1.6M
	Testcase Example:  '[1,1,1,2,2,3]'

	Given an integer array nums sorted in non-decreasing order, remove some
	duplicates in-place such that each unique element appears at most twice. The
	relative order of the elements should be kept the same.

	Since it is impossible to change the length of the array in some languages,
	you must instead have the result be placed in the first part of the array
	nums. More formally, if there are k elements after removing the duplicates,
	then the first k elements of nums should hold the final result. It does not
	matter what you leave beyond the first k elements.

	Return k after placing the final result in the first k slots of nums.

	Do not allocate extra space for another array. You must do this by modifying
	the input array in-place with O(1) extra memory.

	Custom Judge:
	The judge will test your solution with the following code:

	int[] nums = [...]; // Input array
	int[] expectedNums = [...]; // The expected answer with correct length

	int k = removeDuplicates(nums); // Calls your implementation

	assert k == expectedNums.length;
	for (int i = 0; i < k; i++) {
	⁠   assert nums[i] == expectedNums[i];
	}

	If all assertions pass, then your solution will be accepted.

	Example 1:
		Input: nums = [1,1,1,2,2,3]
		Output: 5, nums = [1,1,2,2,3,_]
		Explanation: Your function should return k = 5, with the first five elements
		of nums being 1, 1, 2, 2 and 3 respectively.
		It does not matter what you leave beyond the returned k (hence they are
		underscores).
	Example 2:
		Input: nums = [0,0,1,1,1,1,2,3,3]
		Output: 7, nums = [0,0,1,1,2,3,3,_,_]
		Explanation: Your function should return k = 7, with the first seven
		elements of nums being 0, 0, 1, 1, 2, 3 and 3 respectively.
		It does not matter what you leave beyond the returned k (hence they are
		underscores).
*/
/*
Accepted
167/167 cases passed (0 ms)
Your runtime beats 100 % of golang submissions
Your memory usage beats 7.22 % of golang submissions (3.6 MB)
*/
func (h *LCHandler) RemoveSortedDuplicatesPart2(nums []int) int {
	return removeDuplicatesPart2(nums)
}

// @lc code=start
// NOTE this is a followup to another problem, will need to rename
// to run uncomment below line and comment line 78
// func removeDuplicates(nums []int) int {
func removeDuplicatesPart2(nums []int) int {
	failcaseCount := 0
	collisionMap := make(map[int]int)
	for index, value := range nums {
		hit := collisionMap[value]
		// if we've already come across this number it is a duplicate
		if hit > 1 {
			failcaseCount++
			// go past max value for sorting later and add to count
			nums[index] = 101
		} else {
			collisionMap[value] = hit + 1
		}
	}

	slices.Sort(nums)
	return len(nums) - failcaseCount
}

// @lc code=end
