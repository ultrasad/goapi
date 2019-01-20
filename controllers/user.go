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
func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		//user := new(models.User)
		user := models.CreateUser()
		var err error
		return c.JSON(http.StatusCreated, user)
	}
}

//GetUsers is get user
func GetUsers(c echo.Context) error {
	result := models.GetUser()
	//println("foo")
	return c.JSON(http.StatusOK, result)
}
