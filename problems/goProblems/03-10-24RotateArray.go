package goProblems

import (
	"fmt"
	"strings"
)

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

func (h *Handler) rotateArray(nums []int, k int) string {
	// this seems valid but slow.
	// loop over the array k times each time appending the last element to the front

	for k > 0 {
		toMove := nums[len(nums)-1]
		// after you grab the last element, loop over the array and move each element back one.
		for i := len(nums) - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = toMove
		k--
	}

	return strings.Join(strings.Fields(fmt.Sprint(nums)), ",")

	// this is probably a better way to handle it depending on the size of the array but teh leet code doesn't accept new array returns,
	// only original array modifications. the above code compiles and passes all tests
	// newFront := []int{}
	// // start at the back and traverse the last k elements of the array.
	// // add them in that order to new front
	// for i := len(nums) - 1; i >= len(nums)-k; i-- {
	// 	newFront = append(newFront, nums[i])
	// }

	// // append the new slice of popped elements to the front of the slice of nums[:len-k]
	// toReturn := append(newFront, nums[:len(nums)-k]...)
	// fmt.Println(toReturn)
	// return "array values"
	// return string(currentHighestNumber)
}
