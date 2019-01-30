package models

import (

	//"github.com/jinzhu/gorm"

	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/ultrasad/goapi/db"
	gormdb "github.com/ultrasad/goapi/db/gorm"
)

//BaseModel is default field on table users
type BaseModel struct {
	ID        uint64     `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

//User is user
type (
	User struct {
		//BaseModel
		ID         uint64    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id"`
		Prefix     string    `json:"prefix"`
		Name       string    `json:"name"`
		Email      string    `json:"email"`
		CreateDate string    `json:"create_date, string"`
		Timestamp  time.Time `json:"timestamp" gorm:"column:timestamp" sql:"DEFAULT:current_timestamp"`
		//CreatedAt *time.Time `json:"created_at"`
	}
)

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
	tx := gormdb.DBManager().Begin() // start db transaction
	defer tx.Commit()
	err = fn(tx)
	// close db transaction
	return err
}

/*
func WithinTransaction(fn DBFunc) (err error) {
	tx := gormdb.DBManager().Begin() // start db transaction
	defer tx.Commit()
	err = fn(tx)
	// close db transaction
	return err
}
*/

// FindAll ...
// Helper function to find records by using 'WithinTransaction'
/*
func FindAll(v interface{}) (err error) {
	return WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Find(v).Error; err != nil {
			tx.Rollback() // rollback db transaction
			return err
		}
		return err
	})
}
*/

// Save ...
// Helper function to save gorm model to database by using 'WithinTransaction'
/*
func Save(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if gormdb.DBManager().NewRecord(v) {
			fmt.Println("NewRecord => ", err)
			return err
		}
		if err = tx.Save(v).Error; err != nil {
			fmt.Println("Save => ", err)
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}
*/

// Create ...
// Helper function to insert gorm model to database by using 'WithinTransaction'
func Create(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if !gormdb.DBManager().NewRecord(v) {
			return err
		}
		if err = tx.Create(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

//CreateUserWithTransection is create user with transection
func CreateUserWithTransection(u *User) (*User, error) {
	err := Create(u)
	return u, err
}

//CreateUser is create user
func CreateUser(v interface{}) error {
	return WithinTransaction(func(tx *gorm.DB) (err error) {
		// check new object
		if !gormdb.DBManager().NewRecord(v) {
			return err
		}
		if err = tx.Create(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}
		return err
	})
}

//GetUser is get user
func GetUser(id string) User {
	db := gormdb.ConnectMySQL()
	user := User{}

	//err := db.Debug().Where("name = ?", "Hanajung").Order("id desc, name").Find(&user).Error
	err := db.Debug().Order("id desc, name").Last(&user, id).Error
	if err != nil {
		fmt.Print(err)
	}

	//result.Users = user
	fmt.Println("user => ", user)

	return user
}

//GetUsers is get all user
func GetUsers() Users {
	db := gormdb.ConnectMySQL()
	result := Users{}
	user := []User{}

	//err := db.Debug().Where("name = ?", "Hanajung").Order("id desc, name").Find(&user).Error
	err := db.Debug().Order("id desc, name").Find(&user).Error
	if err != nil {
		fmt.Print(err)
	}

	/*
		if user != nil {
			result.Users = user
			for _, rows := range user {
				fmt.Println("rows => ", rows)
			}
		}
	*/

	result.Users = user
	//fmt.Println("User => ", user)

	return result
}

//GetUserMain is get user
func GetUserMain() Users {
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

//CreateAnimals is create example
func CreateAnimals(db *gorm.DB) error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&User{Name: "Andrew"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&User{Name: "Peter"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
