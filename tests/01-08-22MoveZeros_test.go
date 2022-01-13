package tests

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Given an array nums, write a function to move all 0's to the end
// of it while maintaining the relative order of the non-zero elements.
// You must do this in-place without making a copy of the array.
// Minimize the total number of operations.

// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]
func (h *Handler) MoveZero(logger echo.Context) {
	// success case
	// input := "..R...L..R."
	// fail case
	// input := []int{3, 5, 16, 14, 5, 12}

	returnString := "Pythag tripples for set: "

	logger.HTML(http.StatusOK, returnString)
}
