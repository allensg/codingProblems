package problems

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// You are given a positive integer N which represents the number of steps in a staircase.
// You can either climb 1 or 2 steps at a time.
// Write a function that returns the number of unique ways to climb the stairs.

func (h *Handler) staircase(c echo.Context) {
	input := []int{3, 5, 12, 5, 13}
	marshalled, _ := json.Marshal(input)
	strings.Trim(string(marshalled), "[]")

	returnString := "Pythag tripples for set test "
	returnString = returnString + strings.Trim(string(marshalled), "[]")

	//
	c.HTML(http.StatusOK, returnString)
}
