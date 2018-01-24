package main 

import (
	"fmt" 
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"database/sql"

)

var db *sql.DB

const ( 
	DB_USER     = "yinghua"
	DB_PASSWORD = "123"
	DB_NAME     = "company"
	COMPANY_FILE_DIRECTORY string = "/home/yinghua/Documents/FYP1/FYP-data/company-data/company-data-full.csv"
	ENTRIES	    = 3997101 
	CONCURRENCY = 1000000
)

//====================================================
//function to check error and print error messages
//====================================================
func checkErr(err error, message string) {
	if err != nil {
		panic(message + " err: " + err.Error())
	}
}

//====================================================
// initialize connection with database
//====================================================
func initDB() {

	// create dbString to prevent injection 
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME	)
	sqldb, err := sql.Open("postgres", dbInfo)
	checkErr(err, "Initialize database")
	
	sqldb.SetMaxOpenConns(299)
	db = sqldb
}

func main() { 
	//	importCSV()
	
//	import_uri() 
	import_partnership()
	

}
