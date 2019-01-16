package main

import (
	"net/http"

	"github.com/ultrasad/goapi/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	e.GET("/users", controllers.GetUsers)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
