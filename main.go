package main

import (
	"github.com/labstack/echo"
	"github.com/ultrasad/goapi/db/gorm"
	"github.com/ultrasad/goapi/routers"
)

//Server echo
//var Server *echo.Echo

/*
// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
*/

/*
//CustomContext is custom contact
type CustomContext struct {
	echo.Context
}

//Foo is custom contact
func (c *CustomContext) Foo() {
	println("foo")
}

//Bar is custom contact
func (c *CustomContext) Bar() {
	println("bar")
}
*/

func main() {
	// Echo instance
	//e := echo.New()
	/*
		Server = echo.New()

		// Middleware
		Server.Use(middleware.Logger())
		Server.Use(middleware.Recover())

		fmt.Printf("Echo Starting...")

		// CORS
		Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))

		Server.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "OK")
		})
	*/

	// Router
	//Move to Routers
	/*
		Server.GET("/", hello)
		Server.GET("/users", controllers.GetUsers)
	*/

	/*
		Server.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				cc := &CustomContext{c}
				return h(cc)
			}
		})

		Server.GET("/", func(c echo.Context) error {
			cc := c.(*CustomContext)
			cc.Foo()
			cc.Bar()
			return cc.String(200, "OK")
		})
	*/

	//Echo
	e := echo.New()

	// init database
	gorm.ConnectMySQL()

	// init server
	routers.Init(e)

	// Server
	e.Logger.Fatal(e.Start(":1323"))
}
