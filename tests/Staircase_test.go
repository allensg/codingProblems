package tests

import (
	"fmt"
	"testing"

	"github.com/allensg/codingProblems/problems/goProblems"
)

func TestStaircase(t *testing.T) {
	// so for inputs: 0 | 1 | 2 | 3 | 4 | 5 | 6 |  7 |  8 |  9 | 10 | 11 |  12 |  13 |  14 |...
	// we should see: 0 | 1 | 1 | 2 | 3 | 5 | 8 | 13 | 21 | 34 | 55 | 89 | 144 | 233 | 377 |...

	// Initialize handler
	problems := &goProblems.Handler{}
	var tests = []struct {
		input int
		ans   int
		desc  string
	}{
		{0, 0, "Edge case 0"},
		{1, 1, "Edge case 1"},
		{2, 1, "Edge case 2"},
		{3, 2, "Success case 3"},
		{8, 21, "Success case 8"},
		{14, 377, "Success case 14"},
	}

	for _, testValues := range tests {

		testname := fmt.Sprintf("%s", testValues.desc)
		t.Run(testname, func(t *testing.T) {
			_, outAnswer := problems.Staircase(testValues.input)
			if outAnswer != testValues.ans {
				t.Errorf("input: %d | want %d | got %d", testValues.input, testValues.ans, outAnswer)
			} else {
				t.Logf("input: %d | want %d | got %d", testValues.input, testValues.ans, outAnswer)
			}
		})
	}
}
