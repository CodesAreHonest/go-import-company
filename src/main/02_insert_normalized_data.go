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

func insert_distinct_mortgages() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert company_mortgages data")
	var sqlStatement = "INSERT INTO company_mortgages (com_num_mortchanges, com_num_mortoutstanding, com_num_mortpartsatisfied, com_num_mortsatisfied) VALUES ($1, $2, $3, $4);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert company_partnership")
	
	// 
	for i := len(mortchargesArray); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(mortchargesArray[i], mortoutstandingArray[i], mortpartsatisfiedArray[i], mortsatisfiedArray[i])
			checkErr(err, "Insert partnership statement execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}

func insert_distinct_returns() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert company_returns data")
	var sqlStatement = "INSERT INTO company_returns (com_return_nextduedate, com_return_lastmadeupdate) VALUES ($1, $2);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert company_returns")
	
	// 
	for i := len(nextduedateArray); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(nextduedateArray[i], lastmadeupdateArray[i])
			checkErr(err, "Insert returns statement execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}

func insert_distinct_account_category() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert company_account_category data")
	var sqlStatement = "INSERT INTO company_account_category (com_acc_category) VALUES ($1);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert company_account_category")
	
	// 
	for i := len(accountcategoryArray); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(accountcategoryArray[i])
			checkErr(err, "Insert category execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}

func insert_distinct_account() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert company_account data")
	var sqlStatement = "INSERT INTO company_account (com_acc_refday, com_acc_refmonth, com_acc_nextduedate, com_acc_lastmadeupdate, com_acc_category_id) VALUES ($1, $2, $3, $4, $5);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert company_account")
	
	// 
	for i := len(refdayArray); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(refdayArray[i], refmonthArray[i], a_nextduedateArray[i], a_lastmadeupdateArray[i], categoryIDArray[i])
			checkErr(err, "Insert account execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}

func insert_distinct_conf_statement() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert company_conf_stmt data")
	var sqlStatement = "INSERT INTO company_conf_stmt (com_conf_stmt_nextduedate, com_conf_stmt_lastmadeupdate) VALUES ($1, $2);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert company_conf_stmt")
	
	// 
	for i := len(confstmtnextduedateArray); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(confstmtnextduedateArray[i], confstmtlastmadeupdateArray[i])
			checkErr(err, "Insert conference statement execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}

func insert_distinct_countryoforigin() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert company_countryoforigin data")
	var sqlStatement = "INSERT INTO company_countryoforigin (com_countryoforigin) VALUES ($1);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert company_countryoforigin")
	
	// 
	for i := len(countryoforiginArray); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(countryoforiginArray[i])
			checkErr(err, "Insert country of origin execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}

func insert_distinct_companystatus() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert com_status data")
	var sqlStatement = "INSERT INTO company_status (com_status) VALUES ($1);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert com_status")
	
	for i := len(companystatusArray); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(companystatusArray[i])
			checkErr(err, "Insert company status execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}

func insert_distinct_companycategory() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert com_status data")
	var sqlStatement = "INSERT INTO company_category (com_category) VALUES ($1);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert com_category")
	
	for i := len(companycategoryArray); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(companycategoryArray[i])
			checkErr(err, "Insert company category execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}

func insert_distinct_companysiccode() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert com_status data")
	var sqlStatement = "INSERT INTO company_siccodes (com_siccode1, com_siccode2, com_siccode3, com_siccode4) VALUES ($1, $2, $3, $4);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert com_category")
	
	for i := len(siccode1Array); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(siccode1Array[i], siccode2Array[i], siccode3Array[i], siccode4Array[i])
			checkErr(err, "Insert company category execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}

func insert_distinct_companydetail() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert company_detail data")
	var sqlStatement = "INSERT INTO company_detail (com_name, com_number, com_category_id, com_status_id) VALUES ($1, $2, $3, $4);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert company_detail")
	
	for i := len(companynameArray); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(companynameArray[i], companynumberArray[i], com_category_idArray[i], com_status_idArray[i])
			checkErr(err, "Insert company status execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}

func insert_distinct_previousname() {
	
	sem := make (chan bool, CONCURRENCY) 
	
	fmt.Println("Begin to insert company_previousname data")
	var sqlStatement = "INSERT INTO company_previousname (com_pn1_condate, com_pn1_companyname, com_pn2_condate, com_pn2_companyname, com_pn3_condate, com_pn3_companyname, com_pn4_condate, com_pn4_companyname, com_pn5_condate, com_pn5_companyname, com_pn6_condate, com_pn6_companyname, com_pn7_condate, com_pn7_companyname, com_pn8_condate, com_pn8_companyname, com_pn9_condate, com_pn9_companyname, com_pn10_condate, com_pn10_companyname) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20);"
	
	stmt, err := db.Prepare(sqlStatement)
	checkErr(err, "Prepare insert company_previousname")
	
	for i := len(pn1_condate_Array); i > 0; i-- { 
		sem <- true
		go func () {
			defer func() {<-sem}() 
			_, err := stmt.Exec(pn1_condate_Array[i], pn1_companyname_Array[i], pn2_condate_Array[i], pn2_companyname_Array[i], pn3_condate_Array[i], pn3_companyname_Array[i], pn4_condate_Array[i], pn4_companyname_Array[i], pn5_condate_Array[i], pn5_companyname_Array[i],pn6_condate_Array[i], pn6_companyname_Array[i],pn7_condate_Array[i], pn7_companyname_Array[i], pn8_condate_Array[i], pn8_companyname_Array[i],pn9_condate_Array[i], pn9_companyname_Array[i], pn10_condate_Array[i], pn10_companyname_Array[i])
			checkErr(err, "Insert company previousname execution error") 
		}()
	} 
	
	for i := 0 ; i < cap(sem); i++ { 
		sem <- true
	} 
}





