package problems

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
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
func (h *Handler) Murder(logger echo.Context) {
	// success case
	input, witnesses := []int{3, 6, 3, 4, 1}, []int{}

	for i, k := range input {
		tallest := true
		for _, v := range input[i+1:] {
			if k < v {
				tallest = false
			}
		}
		if tallest {
			witnesses = append(witnesses, i)
		}

	}

	// fail case
	// input := []int{3, 5, 16, 14, 5, 12}
	returnString, count := "", len(witnesses)

	if count > 0 {
		returnString = returnString + fmt.Sprintf("%d people in the crowd (at positions %+q) witnessed the murder.", count, witnesses)
	} else {
		returnString = "There are no witnesses to the murder."
	}

	logger.HTML(http.StatusOK, returnString)
}
