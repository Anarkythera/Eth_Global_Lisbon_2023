package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func ViewConfig(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func EditConfig(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
