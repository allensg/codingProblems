package goProblems

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// Given a list of numbers, find if there exists a pythagorean triplet in that list.
// A pythagorean triplet is 3 variables a, b, c where a2 + b2 = c2
func (h *Handler) PythagTripples(input []int) (string, bool, []int) {
	// success case
	// input = []int{3, 5, 13, 14, 5, 12}
	// fail case
	// input := []int{3, 5, 16, 14, 5, 12}

	// remove duplicate values and sort them
	input = removeDuplicates(input)
	sort.Ints(input)

	// printing stuff
	marshalled, _ := json.Marshal(input)
	returnString := "Pythag tripples for set: " + strings.Trim(string(marshalled), "[]")

	// generate map of squares to be referenced in loop 2
	// key for sum comparison, value for easy printing on return
	squares := make(map[int]int)
	for _, value := range input {
		squares[value*value] = value
	}

	found, a, b, c := false, 0, 0, 0
	// by presorting the array and setting up the sums map
	// we can move in step with only two loops instead of three
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			// check the map of c^2 for an existing instance of a^2 + b^2
			sum := (input[i] * input[i]) + (input[j] * input[j])
			if _, value := squares[sum]; value {
				// h.Logger.Logger().Debug(fmt.Sprintf("i: %d, j: %d, squares[sum]: %d=================sum: %d, ", input[i], input[j], squares[sum], sum))
				a, b, c = input[i], input[j], squares[sum]
				found = value
			}
		}
	}

	if found {
		returnString = returnString + fmt.Sprintf(" | Pythagorean Tripple found at %d^2 + %d^2 = %d^2 ", a, b, c)
	} else {
		returnString = returnString + " | No Pythagorean Tripples found for this set."
	}

	return returnString, found, []int{a, b, c}
}

func removeDuplicates(arr []int) []int {
	keys, toReturn := make(map[int]bool), []int{}

	for _, item := range arr {
		// check the map for an existing instance of that record
		if _, value := keys[item]; !value {
			keys[item] = true
			toReturn = append(toReturn, item)
		}
	}
	return toReturn
}
