package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/allensg/codingProblems/helpers"
	"github.com/allensg/codingProblems/problems"
)

var problem = &problems.Handler{}
var helper = &helpers.Helpers{}

func TestSortLinkedLists(t *testing.T) {
	// Initialize handler

	var tests = []struct {
		input  [][]int
		answer string
		desc   string
	}{
		{[][]int{
			{1, 3, 6},
			{2, 4, 7},
			{5, 8, 9},
		}, "1,2,3,4,5,6,7,8,9", "Given Test Case"},
		// {[]int{8, 5, 4, 3, 2, 1}, 0, "No profit available"},
		// {[]int{1, 2, 3, 4, 5, 44}, 43, "All profit"},
	}

	for _, testValues := range tests {
		fmt.Println("top-------------------------")

		inputArrStr := formatInputArrStr(testValues.input)
		testname := fmt.Sprintf("%s", testValues.desc)
		t.Run(testname, func(t *testing.T) {
			input := GenerateTestDataList(testValues.input)
			returnString, outAnswer := problem.SortLinkedLists(input)
			fmt.Println(returnString)
			displayString := outAnswer.DisplayString()
			if displayString != testValues.answer {
				t.Errorf("input: %s | want %s | got %s", inputArrStr, testValues.answer, displayString)
			} else {
				t.Logf("input: %s | want %s | got %s", inputArrStr, testValues.answer, displayString)
			}
		})
	}

}

func GenerateTestDataList(input [][]int) (listarr []helpers.LinkedList) {

	listarr = []helpers.LinkedList{}
	for _, h := range input {

		list := helpers.LinkedList{}
		for _, cell := range h {
			list.Insert(cell)
		}
		listarr = append(listarr, list)
	}

	return listarr
}

func formatInputArrStr(input [][]int) (returnStr string) {
	inputArrStr := ""
	for _, k := range input {
		currentArrStr := strings.Join(strings.Fields(fmt.Sprint(helper.IntArrToString(k))), ",")
		if inputArrStr != "" {
			inputArrStr = inputArrStr + "," + currentArrStr
		} else {
			inputArrStr = currentArrStr
		}
	}
	return inputArrStr
}
