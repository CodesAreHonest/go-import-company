package main 

import (
	"fmt" 
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"

)

func retrieve_key_from_normalized_table(){
	initDB()
	retrieve_detail_id()
	retrieve_normal_detail()
	retrieve_account_id()
	retrieve_returns_id()
	retrieve_mort_id()
	retrieve_sic_id()
	retrieve_partnership_id()
	retrieve_uri_id()
	retrieve_conf_stmt_id()
	retrieve_previousname_id()
}

func import_company_table() { 
	
	start := time.Now()
	retrieve_key_from_normalized_table()
	insert_company_table()
	fmt.Printf("%.5fs seconds on import company. \n", time.Since(start).Seconds())
}

func insert_company_table() { 
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert company data")
	var sqlStatement = "INSERT INTO company (com_detail_id, com_dissolutiondate, com_incorporationdate, com_countryoforigin, com_careof, com_pobox, com_addressline1, com_addressline2, com_posttown, com_county, com_country, com_postcode, com_acc_id, com_return_id, com_mort_id, com_sic_id, com_partnership_id, com_uri_id, com_pn_id, com_conf_stmt_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert company")
	
	for i := len(dissolutiondateArray); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(com_detail_idArray[i], dissolutiondateArray[i], incorporatedateArray[i], countryoforiginArray[i], careofArray[i], 
				poboxArray[i], addressline1Array[i], addressline2Array[i], posttownArray[i], countyArray[i], 
				countryArray[i], postcodeArray[i], com_acc_idArray[i], com_return_idArray[i], com_mort_idArray[i],
				com_sic_idArray[i], com_partnership_idArray[i], com_uri_idArray[i], com_previousname_idArray[i], com_conf_stmt_idArray[i])
			checkErr(err, "Insert statement execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}
