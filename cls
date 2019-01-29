[1mdiff --git a/controllers/user.go b/controllers/user.go[m
[1mindex 84aa093..76acd15 100644[m
[1m--- a/controllers/user.go[m
[1m+++ b/controllers/user.go[m
[36m@@ -1,7 +1,6 @@[m
 package controllers[m
 [m
 import ([m
[31m-	"encoding/json"[m
 	"fmt"[m
 	"net/http"[m
 [m
[36m@@ -33,38 +32,24 @@[m [mfunc CreateUser() echo.HandlerFunc {[m
 */[m
 [m
 //CreateUser is create new user[m
[31m-func CreateUser(c echo.Context) error {[m
[31m-[m
[31m-	//u := &models.User{}[m
[31m-	//m := echo.Map{}[m
[31m-	//fmt.Println("mm => ", m)[m
[31m-	/*if err := c.Bind(&m); err != nil {[m
[31m-		return err[m
[31m-	}*/[m
[31m-	//return c.JSON(200, m)[m
[31m-[m
[31m-	jsonMap := make(map[string]interface{})[m
[31m-	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)[m
[31m-	if err != nil {[m
[31m-		return err[m
[31m-	}[m
[31m-	//json_map has the JSON Payload decoded into a map[m
[31m-	name := jsonMap["name"][m
[31m-	email := jsonMap["email"][m
[31m-[m
[31m-	u := &models.User{Name: name.(string), Email: email.(string)}[m
[31m-[m
[31m-	result, error := models.CreateUser(u)[m
[31m-	//error := models.Create(c.Bind(&m))[m
[32m+[m[32mfunc CreateUser(c echo.Context) (err error) {[m
[32m+[m	[32m//u := &models.User{Name: "Micle", Email: "Micle@gmail.com"}[m
[32m+[m	[32mu := new(models.User)[m
[32m+[m	[32m//result, error := models.CreateUser(u)[m
[32m+[m	[32merror := models.Create(c.Bind(u))[m
 	if error != nil {[m
[31m-		fmt.Println("error con => ", error)[m
[32m+[m		[32mfmt.Print("error con => ", error)[m
 	}[m
 [m
[31m-	return c.JSON(http.StatusOK, result)[m
[32m+[m	[32m/*[m
[32m+[m		[32mresult := new(models.User)[m
[32m+[m		[32mif err = c.Bind(result); err != nil {[m
[32m+[m			[32mreturn[m
[32m+[m		[32m}[m
[32m+[m	[32m*/[m
 [m
 	//return c.JSON(http.StatusOK, result)[m
[31m-	//return c.JSON(http.StatusOK, m)[m
[31m-[m
[32m+[m	[32mreturn c.JSON(http.StatusOK, u)[m
 }[m
 [m
 //GetUsers is get user[m
[1mdiff --git a/db/gorm/gorm.go b/db/gorm/gorm.go[m
[1mindex 54b577e..4046cfe 100644[m
[1m--- a/db/gorm/gorm.go[m
[1m+++ b/db/gorm/gorm.go[m
[36m@@ -43,7 +43,7 @@[m [mfunc ConnectMySQL() *gorm.DB {[m
 	return db[m
 }[m
 [m
[31m-//DBManager retutn gorm db[m
[32m+[m[32m//DBManager return gorm db[m
 func DBManager() *gorm.DB {[m
 	return db[m
 }[m
[1mdiff --git a/models/user.go b/models/user.go[m
[1mindex aa62cee..2c5ac79 100644[m
[1m--- a/models/user.go[m
[1m+++ b/models/user.go[m
[36m@@ -53,6 +53,16 @@[m [mfunc WithinTransaction(fn DBFunc) (err error) {[m
 	return err[m
 }[m
 [m
[32m+[m[32m/*[m
[32m+[m[32mfunc WithinTransaction(fn DBFunc) (err error) {[m
[32m+[m	[32mtx := gormdb.DBManager().Begin() // start db transaction[m
[32m+[m	[32mdefer tx.Commit()[m
[32m+[m	[32merr = fn(tx)[m
[32m+[m	[32m// close db transaction[m
[32m+[m	[32mreturn err[m
[32m+[m[32m}[m
[32m+[m[32m*/[m
[32m+[m
 // FindAll ...[m
 // Helper function to find records by using 'WithinTransaction'[m
 /*[m
[36m@@ -105,9 +115,10 @@[m [mfunc Create(v interface{}) error {[m
 [m
 //CreateUser is create user[m
 func CreateUser(u *User) (*User, error) {[m
[32m+[m	[32m//func CreateUser(u interface{}) (*User, error) {[m
 	//return Users{}[m
 	//m := User{ID: "3", Name: "Anat", Email: "Anat@gmail.com"}[m
[31m-	//u := &User{Name: "PeterXX", Email: "PeterXX@gmail.com"}[m
[32m+[m	[32m//u := &User{Name: "Peter", Email: "Peter@gmail.com"}[m
 	var err error[m
 	//fmt.Println("m => ", m)[m
 [m
[36m@@ -116,37 +127,13 @@[m [mfunc CreateUser(u *User) (*User, error) {[m
 [m
 	err = Create(u)[m
 [m
[32m+[m	[32m//user := &User{}[m
[32m+[m
 	//fmt.Println("m &m => ", &m)[m
 	fmt.Println("Create m err => ", err)[m
 [m
 	return u, err[m
[31m-}[m
[31m-[m
[31m-//CreateAnimals is create example[m
[31m-func CreateAnimals(db *gorm.DB) error {[m
[31m-	// Note the use of tx as the database handle once you are within a transaction[m
[31m-	tx := db.Begin()[m
[31m-	defer func() {[m
[31m-		if r := recover(); r != nil {[m
[31m-			tx.Rollback()[m
[31m-		}[m
[31m-	}()[m
[31m-[m
[31m-	if err := tx.Error; err != nil {[m
[31m-		return err[m
[31m-	}[m
[31m-[m
[31m-	if err := tx.Create(&User{Name: "Andrew"}).Error; err != nil {[m
[31m-		tx.Rollback()[m
[31m-		return err[m
[31m-	}[m
[31m-[m
[31m-	if err := tx.Create(&User{Name: "Peter"}).Error; err != nil {[m
[31m-		tx.Rollback()[m
[31m-		return err[m
[31m-	}[m
[31m-[m
[31m-	return tx.Commit().Error[m
[32m+[m	[32m//return user, err[m
 }[m
 [m
 //GetUser is get user[m
[36m@@ -287,3 +274,30 @@[m [mfunc GetUserDefault() Users {[m
 	}[m
 	return result[m
 }[m
[32m+[m
[32m+[m[32m//CreateAnimals is create example[m
[32m+[m[32mfunc CreateAnimals(db *gorm.DB) error {[m
[32m+[m	[32m// Note the use of tx as the database handle once you are within a transaction[m
[32m+[m	[32mtx := db.Begin()[m
[32m+[m	[32mdefer func() {[m
[32m+[m		[32mif r := recover(); r != nil {[m
[32m+[m			[32mtx.Rollback()[m
[32m+[m		[32m}[m
[32m+[m	[32m}()[m
[32m+[m
[32m+[m	[32mif err := tx.Error; err != nil {[m
[32m+[m		[32mreturn err[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32mif err := tx.Create(&User{Name: "Andrew"}).Error; err != nil {[m
[32m+[m		[32mtx.Rollback()[m
[32m+[m		[32mreturn err[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32mif err := tx.Create(&User{Name: "Peter"}).Error; err != nil {[m
[32m+[m		[32mtx.Rollback()[m
[32m+[m		[32mreturn err[m
[32m+[m	[32m}[m
[32m+[m
[32m+[m	[32mreturn tx.Commit().Error[m
[32m+[m[32m}[m
