package leetcodeProblems

import (
	"github.com/labstack/echo/v4"
)

type (
	// Handler contains reference to whatever I need specifically for leetcode
	// lots of problems reference earlier problems so this is a way to tie that all
	// together
	LCHandler struct {
		Env    map[string]string
		Logger echo.Context
	}
)
