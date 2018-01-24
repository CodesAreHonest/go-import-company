package main

import (
	"fmt" 
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var ( 
	uriArray [] string 
	genPartnerArray [] string 
	limPartnerArray [] string 
	mortchargesArray [] int 
	mortoutstandingArray [] int 
	mortpartsatisfiedArray [] int 
	mortsatisfiedArray [] int 
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

func retrieve_distinct_partnership() {
	
	initDB()
	
	fmt.Println("Begin to retrieve partnership from company_rawdata")
	rows, err := db.Query("SELECT DISTINCT numgenpartners, numlimpartners FROM company_rawdata;")

	checkErr(err, "Error on query partnership")	
	
	var genpartner string 
	var limpartner string 

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

