package main

import (
	"net/http"
	"os"

	emailProblems "github.com/allensg/codingProblems/problems/goProblems/emailProblems"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.DEBUG)

	envVars := make(map[string]string)

	// Initialize handler
	problems := &emailProblems.Handler{
		Env: envVars,
	}

	e.GET("/", func(c echo.Context) error {
		problems.Logger = c
		// return c.HTML(http.StatusOK, "Hello, Docker! <3")
		// input := []int{3, 5, 13, 14, 5, 12}
		returnString, _ := problems.CountInvalidParens("[}")
		c.HTML(http.StatusOK, returnString)
		return nil
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
