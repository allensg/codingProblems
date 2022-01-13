package problems

import (
	"fmt"
	"strings"
)

// #
// #
// # #
// ####
// ####
// #####
// 36341         x (murder scene)

// There are n people lined up, and each have a height represented as an integer.
// A murder has happened right in front of them, and only people who are taller than
// everyone in front of them are able to see what has happened. How many witnesses are there?

// Input: [3, 6, 3, 4, 1]
// Output: 3 (only 6,4,1 are able to see in front of them)
func (h *Handler) Murder(input []int) (returnString string, witnesses []int) {

	for currentPerson, currentHeight := range input {
		tallest := true
		for _, compareHeight := range input[currentPerson+1:] {
			if currentHeight < compareHeight {
				tallest = false
			}
		}
		if tallest {
			witnesses = append(witnesses, currentPerson)
		}

	}

	count := len(witnesses)
	if count > 0 {
		positions := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(witnesses)), ","), "[]")
		returnString = returnString + fmt.Sprintf("%d people in the crowd (at positions %s) witnessed the murder.", count, positions)
	} else {
		returnString = "There are no witnesses to the murder."
	}

	return returnString, witnesses
}
