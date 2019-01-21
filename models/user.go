package models

import (

	//"github.com/jinzhu/gorm"

	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ultrasad/goapi/db"
	gormdb "github.com/ultrasad/goapi/db/gorm"
)

//User is user
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//Users is user
type Users struct {
	Users []User
}

//DBFunc gorm return error
type (
	DBFunc func(tx *gorm.DB) error // func type which accept *gorm.DB and return error
)

// WithinTransaction ...
// accept DBFunc as parameter
// call DBFunc function within transaction begin, and commit and return error from DBFunc
func WithinTransaction(fn DBFunc) (err error) {
	//tx := cgorm.DBManager().Begin() // start db transaction
	tx := gormdb.DBManager().Begin() // start db transaction
	defer tx.Commit()
	err = fn(tx)
	// close db transaction
	return err
}

// FindAll ...
// Helper function to find records by using 'WithinTransaction'
func FindAll(v interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Find(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}

//CreateUser is create user
func CreateUser() Users {
	return Users{}
}

//GetUserByID is get user
func GetUserByID(id uint64) User {
	db := gormdb.ConnectMySQL()
	user := User{}

	//err := db.Debug().Where("name = ?", "Hanajung").Order("id desc, name").Find(&user).Error
	err := db.Debug().Order("id desc, name").Find(&user).Error
	if err != nil {
		fmt.Print(err)
	}

	//result.Users = user
	fmt.Println("user => ", user)

	return user
}

//GetAllUsers is get all user
func GetAllUsers() Users {
	db := gormdb.ConnectMySQL()
	result := Users{}
	user := []User{}

	//err := db.Debug().Where("name = ?", "Hanajung").Order("id desc, name").Find(&user).Error
	err := db.Debug().Order("id desc, name").Find(&user).Error
	if err != nil {
		fmt.Print(err)
	}

	result.Users = user
	//fmt.Println("User => ", user)

	return result
}

//GetUser is get user
func GetUser() Users {
	db := gormdb.ConnectMySQL()
	result := Users{}
	user := []User{}
	//rows := db.Find(&user)
	//rows := db.Find(&user)
	/*
		if err := db.Find(&user).Error; err != nil {
			//c.String(http.StatusNotFound, "not found page")
			fmt.Println(err)
		}
	*/

	//defer db.Close()
	/*
		if rows, err := db.Debug().Where("name = ?", "Hanajung").Order("id desc, name").Find(&result).Rows(); err != nil {
			fmt.Println(err)
		} else {
			defer rows.Close()
			for rows.Next() {
				//db.ScanRows(rows, &user)
				err2 := db.ScanRows(rows, &user)
				if err2 != nil {
					fmt.Print(err2)
				}

				fmt.Println("users => ", user)
				result.Users = append(result.Users, user)
			}
		}
	*/

	//err := db.Debug().Where("name = ?", "Hanajung").Order("id desc, name").Find(&user).Error
	err := db.Debug().Order("id desc, name").Find(&user).Error
	if err != nil {
		fmt.Print("error db debug => ", err)
	}

	for _, ar := range user {
		fmt.Println(ar.ID)
		result.Users = append(result.Users, ar)
	}

	/*
		for rows.Next() {
			//return result
			err2 := db.ScanRows(rows, &user)
			fmt.Println("rows => ", rows)
			fmt.Println("result => ", result)

			if err2 != nil {
				fmt.Print(err2)
			}
			result.Users = append(result.Users, user)
		}
	*/

	//result.Users = append(result.Users, user)

	//fmt.Println("user => ", user)
	//fmt.Println("result => ", result)

	return result
}

//GetUserDefault is get user
func GetUserDefault() Users {
	con := db.CreateCon()

	sqlStatement := "SELECT id, name, email FROM users order by id"

	rows, err := con.Query(sqlStatement)
	//fmt.Println(rows)
	//fmt.Println(err)
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
