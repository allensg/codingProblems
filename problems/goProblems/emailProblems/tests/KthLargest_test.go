package tests

import (
	"fmt"
	"testing"

	"github.com/allensg/codingProblems/helpers"
	emailProblems "github.com/allensg/codingProblems/problems/goProblems/emailProblems"
)

func TestFindKthLargest(t *testing.T) {
	// Initialize handler
	problems, helper := &emailProblems.Handler{}, &helpers.Helpers{}

	var tests = []struct {
		input      []int
		toFind     int
		kthLargest int
		desc       string
	}{
		{[]int{3, 5, 2, 4, 6, 8}, 3, 5, "Given Test Case"},
		{[]int{8, 34, 18, 3, 21, 1}, 1, 34, "Found Largest"},
		{[]int{7, 6, 2, 41, 5, 4}, 6, 2, "Found Smallest"},
		{[]int{8}, 2, -1, "Out of Bounds"},
	}

	for _, testValues := range tests {

		inputArrStr := helper.IntArrToString(testValues.input)
		testname := fmt.Sprintf("%s", testValues.desc)
		t.Run(testname, func(t *testing.T) {
			_, outAnswer := problems.FindKthLargest(testValues.input, testValues.toFind)
			if outAnswer != testValues.kthLargest {
				t.Errorf("input: %s | want %d | got %d", inputArrStr, testValues.kthLargest, outAnswer)
			} else {
				t.Logf("input: %s | want %d | got %d", inputArrStr, testValues.kthLargest, outAnswer)
			}
		})
	}
}
