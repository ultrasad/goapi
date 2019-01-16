package models

import (
	"database/sql"
	//database sql
	_ "database/sql"
	"fmt"

	"github.com/ultrasad/goapi/db"
)

//User is user
type User struct {
	ID    string
	Name  string
	Email string
}

//Users is user
type Users struct {
	Users []User
}

var con *sql.DB

//CreateUser is create user
func CreateUser() Users {
	result := Users{}
	return result
}

//GetUser is get user
func GetUser() Users {
	con := db.CreateCon()

	sqlStatement := "SELECT id,name, email FROM users order by id"

	rows, err := con.Query(sqlStatement)
	fmt.Println(rows)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	result := Users{}

	for rows.Next() {
		user := User{}

		err2 := rows.Scan(&user.ID, &user.Name, &user.Email)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}
		result.Users = append(result.Users, user)
	}
	return result
}
