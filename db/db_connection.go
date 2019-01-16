package db

import (
	"database/sql"
	"fmt"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

/*CreateCon Create mysql connection*/
func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_restful_api_sample_dev")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}

/*end mysql connection*/
