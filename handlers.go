package handelrs

import (
	"net/http"

	"github.com/labstack/echo"
)

func home(c echo.Context) error {
	//fmt.Fprint("Hello World")
	return c.String(http.StatusOK, "Hello World")
}

//Router is router
func Router(e *echo.Echo) {
	e.GET("/home", home)
}
