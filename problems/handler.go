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

func (h *Handler) FloatArrToString(input []float64) (output string) {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(input)), ","), "[]")
}
