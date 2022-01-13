package problems

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Given a string with the initial condition of dominoes, where:

// . represents that the domino is standing still
// L represents that the domino is falling to the left side
// R represents that the domino is falling to the right side

// Figure out the final position of the dominoes. If there are dominoes that get pushed on
// both ends, the force cancels out and that domino remains upright.

// Input:  ..R...L..R.
// Output: ..RR.LL..RR
func (h *Handler) FallingDominoes(logger echo.Context) {
	// success case
	// input := "..R...L..R."
	// fail case
	// input := []int{3, 5, 16, 14, 5, 12}

	returnString := " "

	logger.HTML(http.StatusOK, returnString)
}
