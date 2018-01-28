package main

import (
	"fmt" 
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func retrieve_detail_id() {	
	fmt.Println("Begin to retrieve company_detail_id from company_detail")
	rows, err := db.Query("SELECT com_detail_id FROM company_detail;")
	checkErr(err, "Error on query com_detaiL_id statement")
	
	var ( 
		com_detail_id int 
	)

    for rows.Next() {
        err = rows.Scan(&com_detail_id)
		checkErr(err, "Retrieve com_detail_id")
		
		com_detail_idArray = append(com_detail_idArray, com_detail_id) 
    }
    
    fmt.Printf("Company detail id: %d \n", len(com_detail_idArray))
    defer rows.Close()
}

func retrieve_normal_detail() { 
	fmt.Println("Begin to retrieve normal detail from company_rawdata")
	rows, err := db.Query("SELECT dissolutiondate, incorporationdate, countryoforigin, careof, pobox, addressline1, addressline2, posttown, county, country, postcode FROM company_rawdata;")
	checkErr(err, "Error on query normal detail statement")
	
	var ( 
		dissolutiondate string 
		incorporationdate string 
		countryoforigin string 
		careof string 
		pobox string 
		addressline1 string 
		addressline2 string 
		posttown string 
		county string 
		country string 
		postcode string 
	)

    for rows.Next() {
        err = rows.Scan(&dissolutiondate, &incorporationdate, &countryoforigin, &careof, &pobox, &addressline1, &addressline2, &posttown, &county, &country, &postcode)
		checkErr(err, "Retrieve company normal detail")
		
		dissolutiondateArray = append(dissolutiondateArray, dissolutiondate)
		incorporatedateArray = append(incorporatedateArray, incorporationdate) 
		countryoforiginArray = append(countryoforiginArray, countryoforigin) 
		careofArray = append(careofArray, careof) 
		poboxArray = append(poboxArray, pobox) 
		addressline1Array = append(addressline1Array, addressline1) 
		addressline2Array = append(addressline2Array, addressline2) 
		posttownArray = append(posttownArray, posttown)  
		countyArray = append(countyArray, county) 
		countryArray = append(countryArray, country) 
		postcodeArray = append(postcodeArray, postcode)  
    }
    
    fmt.Printf("Dissolution date: %d \n", len(dissolutiondateArray))
    fmt.Printf("Incorporationdate: %d \n", len(incorporatedateArray))
    fmt.Printf("Country of origin: %d \n", len(countryoforiginArray))
    fmt.Printf("Careof: %d \n", len(careofArray))
    fmt.Printf("Pobox: %d \n", len(poboxArray))
    fmt.Printf("Address line 1: %d \n", len(addressline1Array))
    fmt.Printf("Address line 2: %d \n", len(addressline2Array))
    fmt.Printf("Post town: %d \n", len(posttownArray))
    fmt.Printf("County: %d \n", len(countyArray))
    fmt.Printf("Country: %d \n", len(countryArray))
    fmt.Printf("Postcode: %d \n", len(postcodeArray))    
    
    defer rows.Close()
}

func retrieve_account_id() {	
	fmt.Println("Begin to retrieve com_acc_id from company_account")
	rows, err := db.Query("SELECT com_acc_id FROM company_account;")
	checkErr(err, "Error on query com_acc_id statement")
	
	var ( 
		com_acc_id int 
	)

    for rows.Next() {
        err = rows.Scan(&com_acc_id)
		checkErr(err, "Retrieve com_detail_id")
		
		com_acc_idArray = append(com_acc_idArray, com_acc_id) 
    }
    fmt.Printf("Company account id: %d \n", len(com_acc_idArray))
    defer rows.Close()
}

func retrieve_returns_id() {	
	fmt.Println("Begin to retrieve com_return_id from company_returns")
	rows, err := db.Query("SELECT com_return_id FROM company_returns AS return JOIN company_rawdata AS raw ON raw.return_nextduedate = return.com_return_nextduedate AND raw.return_lastmadeupdate = return.com_return_lastmadeupdate;")
	checkErr(err, "Error on query com_return_id statement")
	
	var com_return_id int 

    for rows.Next() {
        err = rows.Scan(&com_return_id)
		checkErr(err, "Retrieve com_return_id")
		
		com_return_idArray = append(com_return_idArray, com_return_id) 
    }
    fmt.Printf("Company return id: %d \n", len(com_return_idArray))
    defer rows.Close()
}

func retrieve_mort_id() {	
	fmt.Println("Begin to retrieve com_mort_id from company_mortgages")
	rows, err := db.Query("SELECT com_mort_id FROM company_mortgages AS mort JOIN company_rawdata AS raw ON mort.com_num_mortchanges = raw.nummortcharges AND mort.com_num_mortoutstanding = raw.nummortoutstanding AND mort.com_num_mortpartsatisfied = raw.nummortpartsatisfied AND mort.com_num_mortsatisfied = raw.nummortsatisfied;")
	checkErr(err, "Error on query com_mort_id statement")
	
	var com_mort_id int 

    for rows.Next() {
        err = rows.Scan(&com_mort_id)
		checkErr(err, "Retrieve com_mort_id")
		
		com_mort_idArray = append(com_mort_idArray, com_mort_id) 
    }
    fmt.Printf("Company mort id: %d \n", len(com_mort_idArray))
    defer rows.Close()
}

func retrieve_sic_id() {	
	fmt.Println("Begin to retrieve com_sic_id from company_siccodes")
	rows, err := db.Query("SELECT com_sic_id FROM company_siccodes AS sic JOIN company_rawdata AS raw ON sic.com_siccode1 = raw.siccode1 AND sic.com_siccode2 = raw.siccode2 AND raw.siccode3 = sic.com_siccode3 AND raw.siccode4 = sic.com_siccode4;")
	checkErr(err, "Error on query com_sic_id statement")
	
	var com_sic_id int 

    for rows.Next() {
        err = rows.Scan(&com_sic_id)
		checkErr(err, "Retrieve com_sic_id")
		
		com_sic_idArray = append(com_sic_idArray, com_sic_id) 
    }
    fmt.Printf("Company mort id: %d \n", len(com_sic_idArray))
    defer rows.Close()
}

func retrieve_partnership_id() {	
	fmt.Println("Begin to retrieve com_partnership_id from company_partnership")
	rows, err := db.Query("SELECT com_partnership_id FROM company_partnership AS part JOIN company_rawdata AS raw ON raw.numgenpartners = part.com_num_genpartners AND raw.numlimpartners = part.com_num_limpartners;")
	checkErr(err, "Error on query com_partnership_id statement")
	
	var com_partnership_id int 

    for rows.Next() {
        err = rows.Scan(&com_partnership_id)
		checkErr(err, "Retrieve com_sic_id")
		
		com_partnership_idArray = append(com_partnership_idArray, com_partnership_id) 
    }
    fmt.Printf("Company partnership: %d \n", len(com_partnership_idArray))
    defer rows.Close()
}

func retrieve_uri_id() {	
	fmt.Println("Begin to retrieve com_uri_id from company_uri")
	rows, err := db.Query("SELECT com_uri_id FROM company_uri AS uri JOIN company_rawdata AS raw ON uri.com_uri = raw.uri;")
	checkErr(err, "Error on query com_uri_id statement")
	
	var com_uri_id int 

    for rows.Next() {
        err = rows.Scan(&com_uri_id)
		checkErr(err, "Retrieve com_uri_id")
		
		com_uri_idArray = append(com_uri_idArray, com_uri_id) 
    }
    fmt.Printf("Company uri: %d \n", len(com_uri_idArray))
    defer rows.Close()
}


func retrieve_previousname_id() {	
	fmt.Println("Begin to retrieve com_pn_id from company_previousname")
	rows, err := db.Query("SELECT com_pn_id FROM company_rawdata AS raw JOIN company_previousname AS pn ON raw.pn1_condate = pn.com_pn1_condate AND raw.pn1_companyname = pn.com_pn1_companyname AND raw.pn2_condate = pn.com_pn2_condate AND raw.pn2_companyname = pn.com_pn2_companyname AND raw.pn3_condate = pn.com_pn3_condate AND raw.pn3_companyname = pn.com_pn3_companyname AND raw.pn4_condate = pn.com_pn4_condate AND raw.pn4_companyname = pn.com_pn4_companyname AND raw.pn5_condate = pn.com_pn5_condate AND raw.pn5_companyname = pn.com_pn5_companyname AND raw.pn6_condate = pn.com_pn6_condate AND raw.pn6_companyname = pn.com_pn6_companyname AND raw.pn7_condate = pn.com_pn7_condate AND raw.pn7_companyname = pn.com_pn7_companyname AND raw.pn8_condate = pn.com_pn8_condate AND raw.pn8_companyname = pn.com_pn8_companyname AND raw.pn9_condate = pn.com_pn9_condate AND raw.pn9_companyname = pn.com_pn9_companyname AND raw.pn10_condate = pn.com_pn10_condate;")
	checkErr(err, "Error on query com_pn_id statement")
	
	var com_pn_id int 

    for rows.Next() {
        err = rows.Scan(&com_pn_id)
		checkErr(err, "Retrieve com_pn_id")
		
		com_previousname_idArray = append(com_previousname_idArray, com_pn_id) 
    }
    fmt.Printf("Company previousname: %d \n", len(com_previousname_idArray))
    defer rows.Close()
} 

func retrieve_conf_stmt_id() {	
	fmt.Println("Begin to retrieve com_conf_stmt_id from company_conf_stmt")
	rows, err := db.Query("SELECT com_conf_stmt_id FROM company_conf_stmt AS stmt JOIN company_rawdata AS raw ON stmt.com_conf_stmt_nextduedate = raw.confstmtnextduedate AND stmt.com_conf_stmt_lastmadeupdate = raw.confstmtlastmadeupdate;")
	checkErr(err, "Error on query com_pn_id statement")
	
	var com_conf_stmt_id int 

    for rows.Next() {
        err = rows.Scan(&com_conf_stmt_id)
		checkErr(err, "Retrieve com_conf_stmt_id")
		
		com_conf_stmt_idArray = append(com_conf_stmt_idArray, com_conf_stmt_id) 
    }
    fmt.Printf("Company conference statement: %d \n", len(com_conf_stmt_idArray))
    defer rows.Close()
}





