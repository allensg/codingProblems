package tests

import (
	"testing"
)

envVars := make(map[string]string)

e := echo.New()

e.Use(middleware.Logger())
e.Use(middleware.Recover())
e.Logger.SetLevel(log.DEBUG)

// Initialize handler
problems := &problems.Handler{
	Env: envVars,
}

problems.PythagTripples(e)

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
