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
		input    []float64
		subArray []int
		sum      float64
		desc     string
	}{
		{[]float64{34, -50, 42, 14, -5, 86}, []int{2, 5}, 137, "Given Test Case"},
		{[]float64{34, -50, 42, 14, -5, 86}, []int{2, 5}, 137, "Success Case"},
	}

	for _, testValues := range tests {
		// outSubArray[0] != testValues.subArray[0] ||
		// outSubArray[1] != testValues.subArray[1]
		inputArrStr := problems.FloatArrToString(testValues.input)
		testname := fmt.Sprintf("%s", testValues.desc)
		t.Run(testname, func(t *testing.T) {
			returnString, _, outSum := problems.MaxContigSubArray(testValues.input)
			fmt.Println(returnString)
			if outSum != testValues.sum {
				t.Errorf("input: %s | want %f | got %f", inputArrStr, testValues.sum, outSum)
			} else {
				t.Logf("input: %s | want %f | got %f", inputArrStr, testValues.sum, outSum)
			}
		})
	}
}
