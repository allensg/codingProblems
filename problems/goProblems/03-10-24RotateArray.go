package goProblems

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

func (h *Handler) rotateArray(nums []int, k int) string {
	k = k % len(nums)
	if k == 0 {
		return ""
	}
	toReturn := make([]int, len(nums))
	// get k to end
	toReturn = append(toReturn, nums[len(nums)-k:]...)
	// get start to k
	toReturn = append(toReturn, nums[:len(nums)-k]...)
	// copy
	// stick glue them together backwards
	// bing bang boom
	// copy(nums, toReturn)
	return ""
}
