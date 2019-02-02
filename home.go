package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

//Home Controller
func Home(c echo.Context) error {
	//fmt.Fprint("Hello World")
	return c.String(http.StatusOK, "Hello World")
}
