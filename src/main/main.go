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
	COMPANY_FILE_DIRECTORY string = "/home/yinghua/Documents/FYP1/FYP-data/company-data/company-data.csv"
	ENTRIES	    = 4077979 
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
		DB_USER, DB_PASSWORD, DB_NAME)
	sqldb, err := sql.Open("postgres", dbInfo)
	checkErr(err, "Initialize database")
	
	sqldb.SetMaxOpenConns(299)
	db = sqldb
}

func main() { 
	
//	importCSVtoDB()

//	import_uri() 
//	import_partnership()
//	import_mortgages()
//	import_returns()
//	import_account_category() 
//	import_account() 
//	import_conference_statement()
//	import_companystatus()	
//	import_companycategory()
//	import_companydetail()
//	import_companysiccode() 
//	import_previousname()
	import_company_table()
}
