package tests

import (
	"testing"

	"github.com/allensg/codingProblems/helpers"
	emailProblems "github.com/allensg/codingProblems/problems/goProblems/emailProblems"
)

func TestPythagTripples(t *testing.T) {

	// Initialize handler
	problems, helper := &emailProblems.Handler{}, &helpers.Helpers{}

	var tests = []struct {
		input []int
		found bool
		arr   []int
		desc  string
	}{
		{[]int{3, 5, 13, 14, 5, 12}, true, []int{5, 13, 12}, "Given Test Case"},
		{[]int{3, 5, 16, 14, 5, 12}, false, []int{5, 13, 12}, "Fail case"},
	}

	for _, testValues := range tests {

		inputArrStr := helper.IntArrToString(testValues.input)
		t.Run(testValues.desc, func(t *testing.T) {
			_, outFound, _ := problems.PythagTripples(testValues.input)
			if outFound != testValues.found {
				t.Errorf("input: %s | want %t | got %t", inputArrStr, testValues.found, outFound)
			} else {
				t.Logf("input: %s | want %t | got %t", inputArrStr, testValues.found, outFound)
			}
		})
	}
}
