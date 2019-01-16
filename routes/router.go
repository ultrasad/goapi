package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ultrasad/goapi/controllers"

	"github.com/labstack/echo/middleware"
)

//Init routes
func Init(e *echo.Echo) {

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	fmt.Printf("Starting...")

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	// Secret route
	OtherRouter(e)

	// Router
	e.GET("/users", controllers.GetUsers)
}
