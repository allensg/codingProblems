package problems

import (
	"fmt"
)

// You are given an array of integers. Find the maximum sum of all possible contiguous subarrays of the array.
// Your solution should run in linear time.

// Input: [34, -50, 42, 14, -5, 86]
// Output: subArray:(indexes of the sub array) [2,5] = [42, 14, -5, 86], sum: 137
func (h *Handler) MaxContigSubArray(input []int) (returnString string, subArray []int, sum int) {

	// the idea here is we iterate along the array keeping track of the sum
	// at the current index and the max so far.
	maxSoFar, maxEndingHere := 0, 0
	// lower and upper bounds for the sub array
	subArray = []int{0, 0}
	for i, k := range input {

		// if the sum at i is greater than the value at i another element is contiguous
		if maxEndingHere+k > k {
			maxEndingHere = maxEndingHere + k
		} else {
			// else we've broken our continuity and we
			// update the lower bound of the sub array
			subArray[0] = i
			maxEndingHere = k
		}

		if maxEndingHere >= maxSoFar {
			// update the upper bound of the sub array
			subArray[1] = i
			maxSoFar = maxEndingHere
		}

		fmt.Println(fmt.Sprintf("i: %d, soFar: %d, endingHere: %d", i, maxSoFar, maxEndingHere))
	}

	returnString = fmt.Sprintf("The maximum sub array sum of %d has been found between positions %d(%d) and %d(%d)", maxSoFar, subArray[0], input[subArray[0]], subArray[1], input[subArray[1]])

	return returnString, subArray, maxSoFar
}
