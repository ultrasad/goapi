package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func Router(e *echo.Echo) {
	e.GET("/home", home)
}
