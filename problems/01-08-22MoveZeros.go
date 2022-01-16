package problems

import "fmt"

// Given an array nums, write a function to move all 0's to the end
// of it while maintaining the relative order of the non-zero elements.
// You must do this in-place without making a copy of the array.
// Minimize the total number of operations.

// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]
func (h *Handler) MoveZeros(input []int) (returnString string, answer []int) {

	zeroOffset := 0
	// move all non zero elements to the front of the array
	for _, k := range input {
		if k != 0 {
			input[zeroOffset] = k
			// keep track of how many elements have been moved
			zeroOffset = zeroOffset + 1
		}
	}

	// backfill the last zerOffset elements in the array with zeroes.
	for i := range input[zeroOffset:] {
		input[zeroOffset+i] = 0
	}

	returnString = fmt.Sprintf("The restructured array for this input is %s", h.IntArrToString(input))
	return returnString, input
}
