package main

import (
	"fmt" 
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
    
    fmt.Printf("Company URI: %d \n", len(uriArray))
    defer rows.Close()
}

func retrieve_distinct_partnership() {
	
	initDB()
	
	fmt.Println("Begin to retrieve partnership from company_rawdata")
	rows, err := db.Query("SELECT DISTINCT numgenpartners, numlimpartners FROM company_rawdata;")

	checkErr(err, "Error on query partnership")	
	
	var genpartner int 
	var limpartner int 

    for rows.Next() {
            
            err = rows.Scan(&genpartner, &limpartner)
			checkErr(err, "Retrieve partnership")
			
			genPartnerArray = append(genPartnerArray, genpartner)
			limPartnerArray = append(limPartnerArray, limpartner)
    }
    
    fmt.Printf("General partner: %d \n", len(genPartnerArray))
    fmt.Printf("Limited partner: %d \n", len(limPartnerArray))

    defer rows.Close()
}

func retrieve_distinct_mortgages() {
	
	initDB()
	
	fmt.Println("Begin to retrieve mortgages from company_rawdata")
	rows, err := db.Query("SELECT DISTINCT nummortcharges, nummortoutstanding, nummortpartsatisfied, nummortsatisfied FROM company_rawdata;")

	checkErr(err, "Error on query mortgages")	
	
	var nummortcharges int 
	var nummortoutstanding int 
	var nummortpartsatisfied int 
	var nummortsatisfied int 

    for rows.Next() {
            
            err = rows.Scan(&nummortcharges, &nummortoutstanding, &nummortpartsatisfied, &nummortsatisfied)
			checkErr(err, "Retrieve mortgages")
			
			mortchargesArray = append(mortchargesArray, nummortcharges)
			mortoutstandingArray = append(mortoutstandingArray, nummortoutstanding)
			mortpartsatisfiedArray = append(mortpartsatisfiedArray, nummortpartsatisfied)
			mortsatisfiedArray = append(mortsatisfiedArray, nummortsatisfied)
    }
    
    fmt.Printf("Mort charges: %d \n", len(mortchargesArray))
    fmt.Printf("Mort outstanding: %d \n", len(mortoutstandingArray))
    fmt.Printf("Mort partsatisfied: %d \n", len(mortpartsatisfiedArray))
    fmt.Printf("mort satisfied: %d \n", len(mortsatisfiedArray))

    defer rows.Close()
}

func retrieve_distinct_returns() {
	
	initDB()
	
	fmt.Println("Begin to retrieve returns from company_rawdata")
	rows, err := db.Query("SELECT DISTINCT return_nextduedate, return_lastmadeupdate FROM company_rawdata;")

	checkErr(err, "Error on query returns")	
	
	var return_nextduedate string 
	var return_lastmadeupdate string 

    for rows.Next() {
            
            err = rows.Scan(&return_nextduedate, &return_lastmadeupdate)
			checkErr(err, "Retrieve mortgages")
			
			nextduedateArray = append(nextduedateArray, return_nextduedate)
			lastmadeupdateArray = append(lastmadeupdateArray, return_lastmadeupdate)

    }
    
    fmt.Printf("Return next due date: %d \n", len(nextduedateArray))
    fmt.Printf("Return last made update: %d \n", len(lastmadeupdateArray))


    defer rows.Close()
}

func retrieve_distinct_account_category() {
	
	initDB()
	
	fmt.Println("Begin to retrieve account category from company_rawdata")
	rows, err := db.Query("SELECT DISTINCT accountcategory FROM company_rawdata;")

	checkErr(err, "Error on query company category")	
	
	var category string 

    for rows.Next() {
            
            err = rows.Scan(&category)
			checkErr(err, "Retrieve category")
			
			accountcategoryArray = append(accountcategoryArray, category)
    }
    
    fmt.Printf("Category: %d \n", len(accountcategoryArray))

    defer rows.Close()
}

func retrieve_distinct_account() {
	
	initDB()
	
	fmt.Println("Begin to retrieve account from company_rawdata")
	rows, err := db.Query("SELECT accountingrefday, accountingrefmonth, account_nextduedate, account_lastmadeupdate, com_acc_category_id FROM company_rawdata AS rawdata JOIN company_account_category AS category ON rawdata.accountcategory = category.com_acc_category; ;")

	checkErr(err, "Error on query account")	
	
	var ( 
		refday int64 
		refmonth int64 
		nextduedate string 
		lastmadeupdate string 
		account_category_id int 
	)

    for rows.Next() {
        err = rows.Scan(&refday, &refmonth, &nextduedate, &lastmadeupdate, &account_category_id)
		checkErr(err, "Retrieve account")
		
		refdayArray = append(refdayArray, refday)
		refmonthArray = append(refmonthArray, refmonth)
		a_nextduedateArray = append(a_nextduedateArray, nextduedate)
		a_lastmadeupdateArray = append(a_lastmadeupdateArray, lastmadeupdate)
		categoryIDArray = append(categoryIDArray, account_category_id)
    }
    
    fmt.Printf("Ref day : %d \n", len(refdayArray))
    fmt.Printf("Ref month: %d \n", len(refmonthArray))
    fmt.Printf("Account nextduedate: %d \n", len(a_nextduedateArray))
    fmt.Printf("Account lastmadeupdate: %d \n", len(a_lastmadeupdateArray))
    fmt.Printf("Category ID: %d \n", len(categoryIDArray))

    defer rows.Close()
}

func retrieve_distinct_conf_statement() {
	
	initDB()
	
	fmt.Println("Begin to retrieve conference statement from company_rawdata")
	rows, err := db.Query("SELECT DISTINCT confstmtnextduedate, confstmtlastmadeupdate FROM company_rawdata;")

	checkErr(err, "Error on query conference statement")	
	
	var ( 
		confstmtnextduedate string 
		confstmtlastmadeupdate string 
	)

    for rows.Next() {
        err = rows.Scan(&confstmtnextduedate, &confstmtlastmadeupdate)
		checkErr(err, "Retrieve conference statement")
		
		confstmtnextduedateArray = append(confstmtnextduedateArray, confstmtnextduedate)
		confstmtlastmadeupdateArray = append(confstmtlastmadeupdateArray, confstmtlastmadeupdate)
		    
    }
    
    fmt.Printf("Conference Statement next due date : %d \n", len(confstmtnextduedateArray))
    fmt.Printf("Conference Statement last made update: %d \n", len(confstmtlastmadeupdateArray))

    defer rows.Close()
}

func retrieve_distinct_countryoforigin() {
	
	initDB()
	
	fmt.Println("Begin to retrieve countryoforigin from company_rawdata")
	rows, err := db.Query("SELECT DISTINCT countryoforigin FROM company_rawdata;")

	checkErr(err, "Error on query countryoforigin statement")	
	
	var countryoforigin string 

    for rows.Next() {
        err = rows.Scan(&countryoforigin)
		checkErr(err, "Retrieve country of origin")
		
		countryoforiginArray = append(countryoforiginArray, countryoforigin) 
    }
    
    fmt.Printf("Country of origin: %d \n", len(countryoforiginArray))

    defer rows.Close()
}

func retrieve_distinct_companystatus() {
	
	initDB()
	
	fmt.Println("Begin to retrieve companystatus from company_rawdata")
	rows, err := db.Query("SELECT DISTINCT companystatus FROM company_rawdata;")

	checkErr(err, "Error on query companystatus statement")	
	
	var companystatus string 

    for rows.Next() {
        err = rows.Scan(&companystatus)
		checkErr(err, "Retrieve companystatus")
		
		companystatusArray = append(companystatusArray, companystatus) 
    }
    
    fmt.Printf("Company status: %d \n", len(companystatusArray))

    defer rows.Close()
}

func retrieve_distinct_companycategory() {
	
	initDB()
	
	fmt.Println("Begin to retrieve companycategory from company_rawdata")
	rows, err := db.Query("SELECT DISTINCT companycategory FROM company_rawdata;")

	checkErr(err, "Error on query companystatus statement")	
	
	var companycategory string 

    for rows.Next() {
        err = rows.Scan(&companycategory)
		checkErr(err, "Retrieve companycategory")
		
		companycategoryArray = append(companycategoryArray, companycategory) 
    }
    
    fmt.Printf("Company category: %d \n", len(companycategoryArray))

    defer rows.Close()
}

func retrieve_distinct_companysiccode() { 
	initDB()
	
	fmt.Println("Begin to retrieve siccode from company_rawdata")
	rows, err := db.Query("SELECT DISTINCT siccode1, siccode2, siccode3, siccode4 FROM company_rawdata;")

	checkErr(err, "Error on query siccode statement")	
	
	var (
		siccode1 string 
		siccode2 string 
		siccode3 string 
		siccode4 string 
	)

    for rows.Next() {
        err = rows.Scan(&siccode1, &siccode2, &siccode3, &siccode4)
		checkErr(err, "Retrieve siccode")
		
		siccode1Array = append(siccode1Array, siccode1) 
		siccode2Array = append(siccode2Array, siccode2) 
		siccode3Array = append(siccode3Array, siccode3) 
		siccode4Array = append(siccode4Array, siccode4) 
    }
    
    fmt.Printf("SIC code 1: %d \n", len(siccode1Array))
    fmt.Printf("SIC code 2: %d \n", len(siccode2Array))
    fmt.Printf("SIC code 3: %d \n", len(siccode3Array))
    fmt.Printf("SIC code 4: %d \n", len(siccode4Array))

    defer rows.Close()
}

func retrieve_distinct_companydetail() {
	
	initDB()
	
	fmt.Println("Begin to retrieve companydetail from company_rawdata")
	rows, err := db.Query("SELECT companyname, companynumber, com_category_id, com_status_id FROM company_rawdata AS raw JOIN company_category AS category ON raw.companycategory = category.com_category JOIN company_status AS status ON raw.companystatus = status.com_status;")
	checkErr(err, "Error on query companystatus statement")
	
	var ( 
		companyname string 
		companynumber string 
		com_category_id int 
		com_status_id int 
	)

    for rows.Next() {
        err = rows.Scan(&companyname, &companynumber, &com_category_id, &com_status_id)
		checkErr(err, "Retrieve companystatus")
		
		companynameArray = append(companynameArray, companyname) 
		companynumberArray = append(companynumberArray, companynumber) 
		com_category_idArray = append(com_category_idArray, com_category_id) 
		com_status_idArray = append(com_status_idArray, com_status_id) 
    }
    
    fmt.Printf("Company name: %d \n", len(companynameArray))
    fmt.Printf("Company number: %d \n", len(companynumberArray))
    fmt.Printf("Company category id: %d \n", len(com_category_idArray))
    fmt.Printf("Company status id: %d \n", len(com_status_idArray))
    defer rows.Close()
}

func retrieve_distinct_previousname() {
	
	initDB()
	
	fmt.Println("Begin to retrieve previousdate from company_rawdata")
	rows, err := db.Query("SELECT DISTINCT pn1_condate, pn1_companyname, pn2_condate, pn2_companyname, pn3_condate, pn3_companyname, pn4_condate, pn4_companyname, pn5_condate, pn5_companyname, pn6_condate, pn6_companyname,pn7_condate, pn7_companyname, pn8_condate, pn8_companyname, pn9_condate, pn9_companyname, pn10_condate, pn10_companyname FROM company_rawdata;")
	checkErr(err, "Error on query previousdate statement")
	
	var ( 
		pn1_condate string 
		pn1_companyname string 
		pn2_condate string 
		pn2_companyname string 
		pn3_condate string 
		pn3_companyname string 
		pn4_condate string 
		pn4_companyname string 
		pn5_condate string 
		pn5_companyname string 
		pn6_condate string 
		pn6_companyname string 
		pn7_condate string 
		pn7_companyname string
		pn8_condate string 
		pn8_companyname string 
		pn9_condate string 
		pn9_companyname string 
		pn10_condate string 
		pn10_companyname string 

	)
	
    for rows.Next() {
        err = rows.Scan(&pn1_condate, &pn1_companyname, &pn2_condate, &pn2_companyname, &pn3_condate, &pn3_companyname, &pn4_condate, &pn4_companyname, &pn5_condate, &pn5_companyname, &pn6_condate, &pn6_companyname, &pn7_condate, &pn7_companyname, &pn8_condate, &pn8_companyname, &pn9_condate, &pn9_companyname, &pn10_condate, &pn10_companyname)
		checkErr(err, "Retrieve previousdate")
		
		pn1_condate_Array = append(pn1_condate_Array, pn1_condate)
		pn1_companyname_Array = append(pn1_companyname_Array, pn1_companyname)
		
		pn2_condate_Array = append(pn2_condate_Array, pn2_condate)
		pn2_companyname_Array = append(pn2_companyname_Array, pn2_companyname)
		pn3_condate_Array = append(pn3_condate_Array, pn3_condate)
		pn3_companyname_Array = append(pn3_companyname_Array, pn3_companyname)
		pn4_condate_Array = append(pn4_condate_Array, pn4_condate)
		
		pn4_companyname_Array = append(pn4_companyname_Array, pn4_companyname)
		pn5_condate_Array = append(pn5_condate_Array, pn5_condate)
		pn5_companyname_Array = append(pn5_companyname_Array, pn5_companyname)
		pn6_condate_Array = append(pn6_condate_Array, pn6_condate)
		pn6_companyname_Array = append(pn6_companyname_Array, pn6_companyname)
		
		pn7_condate_Array = append(pn7_condate_Array, pn7_condate)
		pn7_companyname_Array = append(pn7_companyname_Array, pn7_companyname)
		pn8_condate_Array = append(pn8_condate_Array, pn8_condate)
		pn8_companyname_Array = append(pn8_companyname_Array, pn8_companyname)
		pn9_condate_Array = append(pn9_condate_Array, pn9_condate)
		
		pn9_companyname_Array = append(pn9_companyname_Array, pn9_companyname)
		pn10_condate_Array = append(pn10_condate_Array, pn10_condate)
		pn10_companyname_Array = append(pn10_companyname_Array, pn10_companyname)
    }
    
    fmt.Printf("Company change of date 1: %d \n", len(pn1_condate_Array))
    fmt.Printf("Company change name 1: %d \n", len(pn1_companyname_Array))

    defer rows.Close()
}







