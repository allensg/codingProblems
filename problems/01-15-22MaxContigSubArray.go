package problems

import (
	"fmt"
)

// You are given an array of integers. Find the maximum sum of all possible contiguous subarrays of the array.
// Your solution should run in linear time.

// Input: [34, -50, 42, 14, -5, 86]
// Output: subArray:(indexes of the sub array) [2,5] = [42, 14, -5, 86], sum: 137

func (h *Handler) MaxContigSubArray(input []float64) (returnString string, subArray []int, sum float64) {
	maxSoFar, maxEndingHere := 0.0, 0.0
	// lBound, uBound := 0, 0
	for i, k := range input {

		// maxEndingHere = math.Max(k, maxEndingHere+k)
		if maxEndingHere+k >= k {
			fmt.Println(fmt.Sprintf("max ending update at i: %d, maxEndingHere+k: %f, maxEndingHere: %f", i, maxEndingHere+k, maxEndingHere))
			maxEndingHere = maxEndingHere + k
		}
		// maxSoFar = math.Max(maxSoFar, maxEndingHere)
		if maxEndingHere >= maxSoFar {
			fmt.Println(fmt.Sprintf("max soFar update at i: %d, maxEndingHere: %f, maxSoFar: %f", i, maxEndingHere, maxSoFar))
			maxSoFar = maxEndingHere
		}
		fmt.Println(fmt.Sprintf("i: %d, soFar: %f, endingHere: %f", i, maxSoFar, maxEndingHere))
	}

	return returnString, subArray, maxSoFar
}
