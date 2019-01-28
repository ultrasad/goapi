package controllers

import (
	"encoding/json"
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

	//u := &models.User{}
	//m := echo.Map{}
	//fmt.Println("mm => ", m)
	/*if err := c.Bind(&m); err != nil {
		return err
	}*/
	//return c.JSON(200, m)

	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return err
	}
	//json_map has the JSON Payload decoded into a map
	name := jsonMap["name"]
	email := jsonMap["email"]

	u := &models.User{Name: name.(string), Email: email.(string)}

	result, error := models.CreateUser(u)
	//error := models.Create(c.Bind(&m))
	if error != nil {
		fmt.Println("error con => ", error)
	}

	return c.JSON(http.StatusOK, result)

	//return c.JSON(http.StatusOK, result)
	//return c.JSON(http.StatusOK, m)

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
