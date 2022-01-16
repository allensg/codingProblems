package problems

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	// Handler contains reference to whatever I need across problems
	Handler struct {
		Env    map[string]string
		Logger echo.Context
		// IGDB     *igdb.Client
		// Postgres *sql.DB
		// Trie     *patricia.Trie
	}
)

func (h *Handler) IntArrToString(input []int) (output string) {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(input)), ","), "[]")
}

func (h *Handler) AreIntArraysEqual(a []int, b []int) (equal bool) {
	equal = true
	if len(a) != len(b) {
		equal = false
	} else {
		for i, k := range a {
			if k != b[i] {
				equal = false
			}
		}
	}

	return equal
}
