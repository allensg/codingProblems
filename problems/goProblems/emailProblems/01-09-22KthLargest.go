package goProblems

import (
	"fmt"
	"sort"
)

// Given a list, find the k-th largest element in the list

// Input: list = [3, 5, 2, 4, 6, 8], k = 3
// Output: 5
func (h *Handler) FindKthLargest(input []int, toFind int) (returnString string, kthLargest int) {
	// this is a really simple solution and its late
	// Apparently you can solve this with a quicksearch algoritm but its late
	length := len(input)
	if length < toFind {
		return "There aren't enough values to find the Kth largest.", -1
	}
	sort.Ints(input)
	kthLargest = input[length-toFind]

	return fmt.Sprintf("The %dst/rd/th Largest Element is %d", toFind, kthLargest), kthLargest
}
