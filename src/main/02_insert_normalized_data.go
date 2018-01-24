package main

import (
	"fmt" 
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/**  maximum 1 GB of stack can be utilized, or else will crash 
/ 1 goroutine cost 8kb 
  8kb * 125000 = 1000000 KB = 1GB  
**/

func insert_distinct_uri() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert company_uri data")
	var sqlStatement = "INSERT INTO company_uri (com_uri) VALUES ($1);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert company_uri")
	
	// 
	for i := len(uriArray); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(uriArray[i])
			checkErr(err, "Insert statement execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}

func insert_distinct_partnership() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert company_partnership data")
	var sqlStatement = "INSERT INTO company_partnership (com_num_genpartners, com_num_limpartners) VALUES ($1, $2);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert company_partnership")
	
	// 
	for i := len(genPartnerArray); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(genPartnerArray[i], limPartnerArray[i])
			checkErr(err, "Insert statement execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}


