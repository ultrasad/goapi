package controllers

import (
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

//GetUsers is get user
func GetUsers(c echo.Context) error {
	//result := models.GetUser()
	//result := models.FindAll()

	//result := models.GetUserStruct()
	//result := models.GetAllUsers()
	user := models.GetUserByID(1)
	/*for _, ar := range result.Users {
		fmt.Println("range id => ", ar.ID)
		result.Users = append(result.Users, ar)
	}*/

	//fmt.Println("result => ", result.Users)

	//println("foo")
	//return c.JSON(http.StatusOK, result)
	return c.JSON(http.StatusOK, user)
}
