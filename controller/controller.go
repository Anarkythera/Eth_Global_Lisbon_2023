package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

var ControllerPort = ":8082"

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(ControllerPort))

}
