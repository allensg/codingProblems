package tests

import (
	"fmt"
	"testing"

	"github.com/allensg/codingProblems/problems"
)

func TestMaxContigSubArray(t *testing.T) {
	// Initialize handler
	problems := &problems.Handler{}
	var tests = []struct {
		input    []int
		subArray []int
		sum      int
		desc     string
	}{
		{[]int{34, -50, 42, 14, -5, 86}, []int{2, 5}, 137, "Given Test Case"},
		{[]int{69, 42, 14, -900, 1000, 12, 77, 39}, []int{4, 7}, 1128, "Success Case"},
		{[]int{6, 2, 9, 11, 41, 3, 2, -600}, []int{0, 6}, 74, "Frontloaded Case"},
	}

	for _, testValues := range tests {
		inputArrStr := problems.IntArrToString(testValues.input)
		testname := fmt.Sprintf("%s", testValues.desc)
		t.Run(testname, func(t *testing.T) {
			returnString, outSubArray, outSum := problems.MaxContigSubArray(testValues.input)
			fmt.Println(returnString)
			if outSum != testValues.sum ||
				outSubArray[0] != testValues.subArray[0] ||
				outSubArray[1] != testValues.subArray[1] {
				t.Errorf("input: %s | want %d | got %d", inputArrStr, testValues.sum, outSum)
			} else {
				t.Logf("input: %s | want %d | got %d", inputArrStr, testValues.sum, outSum)
			}
		})
	}
}
