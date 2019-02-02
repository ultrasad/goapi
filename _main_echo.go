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

/*
type addressBook struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Code      int    `json:"code"`
	Phone     string `json:"phone"`
}

//getAddressBookAll is get address book all services
func getAddressBookAll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("get address book all")

	addBook := addressBook{
		Firstname: "Chaiyarin",
		Lastname:  "Niamsuwan",
		Code:      1993,
		Phone:     "0870940955",
	}
	json.NewEncoder(w).Encode(addBook)
}

//handleRequestAddress is call services method
func handleRequestAddress() {
	fmt.Println("Handle Request Address")
	http.HandleFunc("/getAddress", getAddressBookAll)
	http.ListenAndServe(":1323", nil)
}
*/

func mainEcho() {
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
