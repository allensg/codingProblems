package goProblems

import (
	"fmt"
	"sort"
)

// You are given an array of integers. Return the largest product that can
// be made by multiplying any 3 integers in the array.

// solves in O(nlogn)
func (h *Handler) MaxProdOf3(input []int) (returnString string, maxProduct int) {
	maxProduct = 0
	if len(input) < 3 {
		return "Not enough Elements in the array", -1
	} else {
		sort.Ints(input)

		// if the first element is less than 0 we have negatives
		if 0 > input[0] {
			// if there are negatives its the smallest 2 and the biggest
			maxProduct = input[0] * input[1] * input[len(input)-1]
		} else {
			// if no negatives its just the last 3 values in the array
			maxProduct = input[len(input)-3] * input[len(input)-2] * input[len(input)-1]
		}
	}

	returnString = fmt.Sprintf("The Maximum product of three numbers in the input list is %d: ", maxProduct)

	return returnString, maxProduct
}
