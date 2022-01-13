package tests

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// You are given a 2D array of integers.
// Print out the clockwise spiral traversal of the matrix.

// Input: [[1,  2,  3,  4,  5],
// 		[6,  7,  8,  9,  10],
// 		[11, 12, 13, 14, 15],
// 		[16, 17, 18, 19, 20]]
// Output: 1, 2, 3, 4, 5, 10, 15, 20, 19, 18, 17, 16, 11, 6, 7, 8, 9, 14, 13, 12
func (h *Handler) MatrixSpiralPrint(logger echo.Context) {
	// success case
	// input := "..R...L..R."
	// fail case
	// input := []int{3, 5, 16, 14, 5, 12}

	returnString := "Pythag tripples for set: "

	logger.HTML(http.StatusOK, returnString)
}
