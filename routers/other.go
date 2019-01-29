package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

//OtherRouter is other routes
func OtherRouter(e *echo.Echo) {
	//log.Print("Secret")

	e.GET("/secret", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ช้างน้อย ฮานะจัง")
	})
}
