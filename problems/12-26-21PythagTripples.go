package problems

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
)

// Given a list of numbers, find if there exists a pythagorean triplet in that list.
// A pythagorean triplet is 3 variables a, b, c where a2 + b2 = c2
func (h *Handler) PythagTripples(logger echo.Context) {
	// success case
	input := []int{3, 5, 13, 14, 5, 12}
	// fail case
	// input := []int{3, 5, 16, 14, 5, 12}

	// remove duplicate values
	input = removeDuplicates(input)
	// sort
	sort.Ints(input)
	marshalled, _ := json.Marshal(input)
	returnString := "Pythag tripples for set: "
	returnString = returnString + strings.Trim(string(marshalled), "[]")
	found := false
	a, b, c := 0, 0, 0

	for i := 0; i < len(input)-2; i++ {
		for j := i + 1; j < len(input)-1; j++ {
			for k := i + 2; k < len(input); k++ {
				left := (input[i] * input[i]) + (input[j] * input[j])
				right := (input[k] * input[k])
				// logger.Logger().Debug(fmt.Sprintf("i: %d, j: %d, k: %d=================left: %d, right: %d", input[i], input[j], input[k], left, right))
				if left == right {
					a, b, c = input[i], input[j], input[k]
					found = true
					// logger.Logger().Debug("inside if statemetn ------------------------------------------")
				}
			}
		}
	}

	if found {
		// add found message to return string
		returnString = returnString + fmt.Sprintf(" | Pythagorean Tripple found at %d^2 + %d^2 = %d^2 ", a, b, c)
	} else {
		// add not found message to return string
		returnString = returnString + fmt.Sprintf(" | No Pythagorean Tripples found for this set.")
	}

	logger.HTML(http.StatusOK, returnString)
}

func removeDuplicates(arr []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, item := range arr {
		if _, value := keys[item]; !value {
			keys[item] = true
			list = append(list, item)
		}
	}
	return list
}
