package tests

import (
	"testing"
	// problems "github.com/allensg/codingProblems/tests"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {

		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}

}

// func TestPythagTripples(t *testing.T) {
// 	problems := &problems.Handler{
// 		// Env: envVars,
// 	}

// 	answer := problems.PythagTripples()
// 	t.Errorf
// }
