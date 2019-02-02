package handlers

import (
	"github.com/labstack/echo"
)

//Router is router
func Router(e *echo.Echo) {
	e.GET("/home", Home)
}
