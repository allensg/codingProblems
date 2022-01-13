package problems

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Given a list, find the k-th largest element in the list

// Input: list = [3, 5, 2, 4, 6, 8], k = 3
// Output: 5
func (h *Handler) FindKthLargest(logger echo.Context) {
	// success case
	// input := "..R...L..R."
	// fail case
	// input := []int{3, 5, 16, 14, 5, 12}

	returnString := "Pythag tripples for set: "

	logger.HTML(http.StatusOK, returnString)
}
