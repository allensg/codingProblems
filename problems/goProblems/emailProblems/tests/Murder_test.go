package tests

import (
	"fmt"
	"testing"

	"github.com/allensg/codingProblems/helpers"
	emailProblems "github.com/allensg/codingProblems/problems/goProblems/emailProblems"
)

func TestMurder(t *testing.T) {
	// Initialize handler
	problems, helper := &emailProblems.Handler{}, &helpers.Helpers{}

	var tests = []struct {
		input     []int
		count     int
		witnesses []int
		desc      string
	}{
		{[]int{3, 6, 3, 4, 1}, 3, []int{1, 3, 4}, "Given Test Case"},
		{[]int{}, 0, []int{}, "EmptyArr"},
		{[]int{3, 6, 3, 4, 9}, 1, []int{4}, "Tall Guy Blocker Case"},
	}

	for _, testValues := range tests {

		inputArrStr := helper.IntArrToString(testValues.input)
		testname := fmt.Sprintf("%s", testValues.desc)
		t.Run(testname, func(t *testing.T) {
			_, outAnswer := problems.Murder(testValues.input)
			if len(outAnswer) != testValues.count {
				t.Errorf("input: %s | want %d | got %d", inputArrStr, testValues.count, len(outAnswer))
			} else {
				t.Logf("input: %s | want %d | got %d", inputArrStr, testValues.count, len(outAnswer))
			}
		})
	}
}
