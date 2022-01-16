package tests

import (
	"fmt"
	"testing"

	"github.com/allensg/codingProblems/problems"
)

func TestSortLinkedLists(t *testing.T) {
	// Initialize handler
	problems := &problems.Handler{}
	var tests = []struct {
		input  []int
		answer int
		desc   string
	}{
		{[]int{9, 11, 8, 5, 7, 10}, 5, "Given Test Case"},
		{[]int{8, 5, 4, 3, 2, 1}, 0, "No profit available"},
		{[]int{1, 2, 3, 4, 5, 44}, 43, "All profit"},
	}

	for _, testValues := range tests {

		inputArrStr := problems.IntArrToString(testValues.input)
		testname := fmt.Sprintf("%s", testValues.desc)
		t.Run(testname, func(t *testing.T) {
			returnString, outAnswer := problems.BuyAndSell(testValues.input)
			fmt.Println(returnString)
			if outAnswer != testValues.answer {
				t.Errorf("input: %s | want %d | got %d", inputArrStr, testValues.answer, outAnswer)
			} else {
				t.Logf("input: %s | want %d | got %d", inputArrStr, testValues.answer, outAnswer)
			}
		})
	}
}
