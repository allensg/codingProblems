package problems

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// You are given an array of integers. Return the largest product that can
// be made by multiplying any 3 integers in the array.

// Input: [-4, -4, 2, 8]
// Output: 128
func (h *Handler) MaxProdOf3(logger echo.Context) {
	// success case
	// input := "..R...L..R."
	// fail case
	// input := []int{3, 5, 16, 14, 5, 12}

	returnString := "Pythag tripples for set: "

	logger.HTML(http.StatusOK, returnString)
}
