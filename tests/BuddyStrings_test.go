package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/allensg/codingProblems/problems/goProblems"
)

func TestBuddyStrings(t *testing.T) {
	// Initialize handler
	problems := &goProblems.Handler{}

	var tests = []struct {
		input  []string
		answer bool
		desc   string
	}{
		{[]string{"ab", "ba"}, true, "Given Test Case 1"},
		{[]string{"ab", "ab"}, false, "Given Test Case 2"},
		{[]string{"aa", "aa"}, true, "Given Test Case 3"},
		{[]string{"aaaaaaabc", "aaaaaaacb"}, true, "Given Test Case 4"},
		{[]string{"", "aa"}, true, "Given Test Case 5"},
	}

	for _, testValues := range tests {

		inputArrStr := strings.Join(testValues.input, ", ")
		testname := fmt.Sprintf("%s", testValues.desc)
		t.Run(testname, func(t *testing.T) {
			_, outAnswer := problems.BuddyStrings(testValues.input)
			if outAnswer != testValues.answer {
				t.Errorf("input: %s | want %t | got %t", inputArrStr, testValues.answer, outAnswer)
			} else {
				t.Logf("input: %s | want %t | got %t", inputArrStr, testValues.answer, outAnswer)
			}
		})
	}
}
