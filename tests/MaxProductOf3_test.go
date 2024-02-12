package tests

import (
	"fmt"
	"testing"

	"github.com/allensg/codingProblems/helpers"
	"github.com/allensg/codingProblems/problems/goProblems"
)

func TestMaxProdOf3(t *testing.T) {
	// Initialize handler
	problems, helper := &goProblems.Handler{}, &helpers.Helpers{}

	var tests = []struct {
		input  []int
		maxSum int
		desc   string
	}{
		{[]int{-4, -4, 2, 8}, 128, "Given Test Case"},
		{[]int{-12, -6, 0, 4}, 288, "Zero case"},
		{[]int{3, 2}, -1, "Array too small"},
	}

	for _, testValues := range tests {

		inputArrStr := helper.IntArrToString(testValues.input)
		testname := fmt.Sprintf("%s", testValues.desc)
		t.Run(testname, func(t *testing.T) {
			_, outAnswer := problems.MaxProdOf3(testValues.input)
			if outAnswer != testValues.maxSum {
				t.Errorf("input: %s | want %d | got %d", inputArrStr, testValues.maxSum, outAnswer)
			} else {
				t.Logf("input: %s | want %d | got %d", inputArrStr, testValues.maxSum, outAnswer)
			}
		})
	}
}
