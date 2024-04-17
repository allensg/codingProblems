package leetcodeProblems

import "slices"

/*
	@lc app=leetcode id=26 lang=golang

	[26] Remove Duplicates from Sorted Array

	https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/

	algorithms
	Easy (55.48%)
	Likes:    14071
	Dislikes: 18278
	Total Accepted:    4.2M
	Total Submissions: 7.6M
	Testcase Example:  '[1,1,2]'

	Given an integer array nums sorted in non-decreasing order, remove the
	duplicates in-place such that each unique element appears only once. The
	relative order of the elements should be kept the same. Then return the
	number of unique elements in nums.

	Consider the number of unique elements of nums to be k, to get accepted, you
	need to do the following things:


	Change the array nums such that the first k elements of nums contain the
	unique elements in the order they were present in nums initially. The
	remaining elements of nums are not important as well as the size of
	nums.
	Return k.


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
		Input: nums = [1,1,2]
		Output: 2, nums = [1,2,_]
		Explanation: Your function should return k = 2, with the first two elements
		of nums being 1 and 2 respectively.
		It does not matter what you leave beyond the returned k (hence they are
		underscores).
	Example 2:
		Input: nums = [0,0,1,1,1,2,2,3,3,4]
		Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
		Explanation: Your function should return k = 5, with the first five elements
		of nums being 0, 1, 2, 3, and 4 respectively.
		It does not matter what you leave beyond the returned k (hence they are
		underscores).

*/
/*
Accepted
362/362 cases passed (10 ms)
Your runtime beats 15.83 % of golang submissions
Your memory usage beats 9.75 % of golang submissions (4.8 MB)
*/
func (h *LCHandler) RemoveSortedDuplicatesPart1(nums []int) int {
	return removeDuplicatesPart1(nums)
}

// @lc code=start
// NOTE this problem has a followup, will need to rename
// to run uncomment below line and comment line 83
// func removeDuplicates(nums []int) int {
func removeDuplicatesPart1(nums []int) int {
	collisionMap := make(map[int]bool)
	for index, value := range nums {
		hit := collisionMap[value]
		// if we've already come across this number it is a duplicate
		if hit {
			// go past max value for sorting later and add to count
			nums[index] = 101
		} else {
			collisionMap[value] = true
		}
	}

	slices.Sort(nums)
	return len(collisionMap)
}

// @lc code=end
