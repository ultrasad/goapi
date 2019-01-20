package gorm

import (
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

//ConnectGORM is connect with gorm
func ConnectGORM() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(0.0.0.0:3306)"
	DBNAME := "golang_restful_api_sample_dev"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	/*db,err := gorm.Open(DBMS, CONNECT)
	    if err != nil {
	        panic(err.Error())
		}*/

	if dbxxx, err = gorm.Open(DBMS, CONNECT); err != nil {
		panic(err.Error())
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	//db.LogMode(true)

	return db
}
