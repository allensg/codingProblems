package tests

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Example:
// {
//   'CSC300': ['CSC100', 'CSC200'],
//   'CSC200': ['CSC100'],
//   'CSC100': []
// }

// This input should return the order that we need to take these courses:
// ['CSC100', 'CSC200', 'CSCS300']

// You are given a hash table where the key is a course code, and the value is a
// list of all the course codes that are prerequisites for the key.
// Return a valid ordering in which we can complete the courses.
// If no such ordering exists, return NULL
func (h *Handler) CoursesToTake(logger echo.Context) {
	// success case
	// input := "..R...L..R."
	// fail case
	// input := []int{3, 5, 16, 14, 5, 12}

	returnString := "Pythag tripples for set: "

	logger.HTML(http.StatusOK, returnString)
}
