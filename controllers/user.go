package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ultrasad/goapi/models"

	"github.com/labstack/echo"
)

//CreateUser is create new user
func CreateUser(c echo.Context) (err error) {

	jsonMap := make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return err
	}

	//jsonMap has the JSON Payload decoded into a map
	prefix := jsonMap["prefix"]
	name := jsonMap["name"]
	email := jsonMap["email"]
	inputdate := jsonMap["create_date"]
<<<<<<< HEAD
	//createDate, _ := time.Parse("2016-01-02 15:04:05", inputdate.(string))
	//inputdate = fmt.Sprintf("%v 00:00:00", inputdate)
	//datestamp, _ := time.Parse("2006-01-02 15:04:05", inputdate.(string))
	//datestamp.Format("2016-01-02 00:00:00")

	//fmt.Println("datestamp => ", datestamp, ", => ", err)
	fmt.Println("datestamp => ", inputdate)
=======

	//inputdate = fmt.Sprintf("%v 00:00:00", inputdate)
	//datestamp, _ := time.Parse("2006-01-02 15:04:05", inputdate.(string))

	//fmt.Println("inputdate => ", inputdate)
>>>>>>> ae0ad7554be3eecde06a815508ab57a0bb3a153e

	inputtamp := jsonMap["timestamp"]

	//"2006-01-02 15:04:05" is standard format datetime golang
	timestamp, _ := time.Parse("2006-01-02 15:04:05", inputtamp.(string))

	//fmt.Println("timestamp => ", timestamp)

	user := &models.User{Name: name.(string), Email: email.(string), Prefix: prefix.(string), Timestamp: timestamp, CreateDate: inputdate.(string)}
	/*
		result, err := models.CreateUserWithTransection(user)
		//fix format timestamp
		result.Timestamp.Format("2006-01-02 15:04:05")
		if err != nil {
			return
		}
		return c.JSON(http.StatusOK, result)
	*/

	//user.Timestamp.Format("2006-01-02 15:04:05")
	err = models.CreateUser(user)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, user)

	//Bind Data with Process
	/*
		usr := new(models.User)
		if err = c.Bind(usr); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, usr)
	*/
}

//GetUsers is get user
func GetUsers(c echo.Context) error {

	result := models.GetUsers()
<<<<<<< HEAD
	//result := models.GetUserByID(1)

	for _, ar := range result.Users {
		fmt.Println("row => ", ar.Name, ar.CreateDate, ar.Timestamp.Format("2006-01-02 15:04:05"))
		//result.Users = append(result.Users, ar)
	}
=======
>>>>>>> ae0ad7554be3eecde06a815508ab57a0bb3a153e

	for i, ar := range result.Users {
		createDate, _ := time.Parse("2006-01-02T00:00:00Z", ar.CreateDate)
		result.Users[i].CreateDate = createDate.Format("2006-01-02")
	}

	return c.JSON(http.StatusOK, result)
}

//GetUser is get user by id
func GetUser(c echo.Context) error {
	id := c.Param("id")
	result := models.GetUser(id)
	return c.JSON(http.StatusOK, result)
}
