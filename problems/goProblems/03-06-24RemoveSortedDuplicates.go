package goProblems

import (
	"golang.org/x/exp/slices"
)

// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially.
// The remaining elements of nums are not important as well as the size of nums.
// Return k.
func (h *Handler) removeSortedDuplicatesPart1(nums []int, val int) string {
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
	return string(len(collisionMap))
}

// the next problem in the sequence is an expansion on the first. Scoring the second solution and qeustion here.

// Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears AT MOST TWICE.
// The relative order of the elements should be kept the same.

// Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first
// part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

// Return k after placing the final result in the first k slots of nums.

func (h *Handler) removeSortedDuplicatesPart2(nums []int, val int) string {
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
	return string(len(collisionMap))
}
