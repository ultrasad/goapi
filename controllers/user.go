package controllers

import (
	"fmt"
	"net/http"

	"github.com/ultrasad/goapi/models"

	"github.com/labstack/echo"
)

//CreateUser is create user
/*
func CreateUser(c echo.Context) error {
	//user := new(models.User)
	user := models.CreateUser()
	var err error
	return c.JSON(http.StatusCreated, user)
}
*/

//CreateUser is create user
/*
func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		//user := new(models.User)
		user := models.CreateUser()
		var err error
		return c.JSON(http.StatusCreated, user)
	}
}
*/

//CreateUser is create new user
func CreateUser(c echo.Context) error {
	result, error := models.CreateUser()
	if error != nil {
		fmt.Print("error con => ", error)
	}
	return c.JSON(http.StatusOK, result)
}

//GetUsers is get user
func GetUsers(c echo.Context) error {
	//result := models.GetUser()
	//result := models.FindAll()

	//q := new([]int)
	//fmt.Println(*q)
	//fmt.Printf("Value = nil ? =>%t\n", *q == nil)

	//result := models.GetUserStruct()
	result := models.GetUsers()
	//result := models.GetUserByID(1)
	for _, ar := range result.Users {
		fmt.Println("range id => ", ar.Name, ar.Timestamp.Format("2006-01-02 15:04:05"))
		//result.Users = append(result.Users, ar)
	}

	//fmt.Println("result users => ", result.Users)

	//println("foo")
	//return c.JSON(http.StatusOK, result)
	return c.JSON(http.StatusOK, result)
}

//GetUser is get user by id
func GetUser(c echo.Context) error {
	id := c.Param("id")
	result := models.GetUser(id)
	return c.JSON(http.StatusOK, result)
}
