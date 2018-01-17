package main

import (
	"fmt" 
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var ( 
	uriArray [] string 
)

func retrieve_distinct_uri() {
	
	initDB()
	
	fmt.Println("Begin to retrieve uri from company_rawdata")
	rows, err := db.Query("SELECT DISTINCT uri FROM company_rawdata;")

	checkErr(err, "Error on query uri")	
	
	var uri string 

    for rows.Next() {
            
            err = rows.Scan(&uri)
			checkErr(err, "Retrieve uri")
			
			uriArray = append(uriArray, uri)
//			fmt.Println(uri)
    }
    
    fmt.Printf("Postcode detail: %d \n", len(uriArray))
    defer rows.Close()
    
}

