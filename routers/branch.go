package routers

import (
	"net/http"

	"github.com/labstack/echo"
)

func branch(e *echo.Echo) {

	e.GET("/branch", func(c echo.Context) error {
		jsonData := map[string]string{"firstname": "Bundit", "lastname": "Parameesee"}
		return c.JSON(http.StatusOK, jsonData)
	})

}
