package goProblems

import (
	"golang.org/x/exp/slices"
)

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
// The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

// Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
// The remaining elements of nums are not important as well as the size of nums.
// Return k.

// this is kind of cheeky becasuse the parameters of the leetcode problem dictate that no value will be higher than 100.
// so replacing and sorting solves the problem quickly but some people could have a problem with it.
func (h *Handler) removeElement(nums []int, val int) string {
	count := 0
	for index, value := range nums {
		if value == val {
			nums[index] = 101
			count++
		}
	}

	slices.Sort(nums)
	return string(len(nums) - count)
}
