package goProblems

import (
	"github.com/labstack/echo/v4"
)

type (
	// Handler contains reference to whatever I need across email
	Handler struct {
		Env    map[string]string
		Logger echo.Context
		// IGDB     *igdb.Client
		// Postgres *sql.DB
		// Trie     *patricia.Trie
	}
)
