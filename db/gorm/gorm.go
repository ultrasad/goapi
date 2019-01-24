package gorm

import (
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	//mysql dialects
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

//ConnectMySQL is connect mysql db with gorm
func ConnectMySQL() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(0.0.0.0:3306)"
	DBNAME := "golang_restful_api_sample_dev"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True"
	/*db,err := gorm.Open(DBMS, CONNECT)
	    if err != nil {
	        panic(err.Error())
		}*/

	if db, err = gorm.Open(DBMS, CONNECT); err != nil {
		panic(err.Error())
	}

	//defer db.Close()
	// make sure connection is available
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	//db.LogMode(true)

	return db
}

//DBManager retutn gorm db
func DBManager() *gorm.DB {
	return db
}
