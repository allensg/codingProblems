package tests

import (
	"fmt"
	"testing"

	"github.com/allensg/codingProblems/problems"
)

func TestMoveZeross(t *testing.T) {
	// Initialize handler
	problems := &problems.Handler{}
	var tests = []struct {
		input  []int
		answer []int
		desc   string
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}, "Given Test Case"},
		{[]int{8, 0, 0, 3, 0, 1}, []int{8, 3, 1, 0, 0, 0}, "Success Case"},
		{[]int{8, 5, 6, 3, 7, 1}, []int{8, 5, 6, 3, 7, 1}, "No Zeros case"},
	}

	for _, testValues := range tests {

		inputArrStr := problems.IntArrToString(testValues.input)
		testname := fmt.Sprintf("%s", testValues.desc)
		t.Run(testname, func(t *testing.T) {
			returnString, outAnswer := problems.MoveZeros(testValues.input)
			fmt.Println(returnString)
			if !problems.AreIntArraysEqual(testValues.answer, outAnswer) {
				t.Errorf("input: %s | want %d | got %d", inputArrStr, testValues.answer, outAnswer)
			} else {
				t.Logf("input: %s | want %d | got %d", inputArrStr, testValues.answer, outAnswer)
			}
		})
	}
}
