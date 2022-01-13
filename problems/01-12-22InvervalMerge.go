package problems

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// You are given an array of intervals - that is, an array of tuples (start, end).
// The array may not be sorted, and could contain overlapping intervals.
// Return another array where the overlapping intervals are merged.

// Input: [(1, 3), (5, 8), (4, 10), (20, 25)]
// Output: [(1, 3), (4, 10), (20, 25)]
func (h *Handler) MergeIntervals(logger echo.Context) {
	// success case
	// input := "..R...L..R."
	// fail case
	// input := []int{3, 5, 16, 14, 5, 12}

	returnString := "Pythag tripples for set: "

	logger.HTML(http.StatusOK, returnString)
}
