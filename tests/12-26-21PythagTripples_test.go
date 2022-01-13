package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/allensg/codingProblems/problems"
)

func TestPythagTripples(t *testing.T) {

	// Initialize handler
	problems := &problems.Handler{}
	var tests = []struct {
		input []int
		found bool
		arr   []int
		desc  string
	}{
		{[]int{3, 5, 13, 14, 5, 12}, true, []int{5, 13, 12}, "Success case"},
		{[]int{3, 5, 16, 14, 5, 12}, false, []int{5, 13, 12}, "Fail case"},
	}

	for _, testValues := range tests {

		inputArr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(testValues.input)), ","), "[]")
		testname := fmt.Sprintf("%s: %s", testValues.desc, inputArr)
		t.Run(testname, func(t *testing.T) {
			_, outFound, _ := problems.PythagTripples(testValues.input)
			if outFound != testValues.found {
				t.Errorf("got %t, want %t", outFound, testValues.found)
			} else {
				t.Logf("got %t, want %t", outFound, testValues.found)
			}
		})
	}
}
