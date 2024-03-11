package goProblems

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
func (h *Handler) majorityElement(nums []int) string {
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
			return string(value)
		}
		newCount := currentCount + 1
		if newCount >= counts[currentHighestNumber] {
			currentHighestNumber = value
		}
		counts[value] = currentCount + 1
	}

	return string(currentHighestNumber)
}
