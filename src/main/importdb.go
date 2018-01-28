package main

import (
	"fmt" 
)

func importDB() { 
	
	sem := make (chan bool, 400000) 

	
	initDB() 
	fmt.Println("Prepare to import data") 

	var sStmt string = "INSERT INTO company_rawdata1 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55);"

	stmt, err := db.Prepare(sStmt)
	checkErr(err, "Prepare insert com_category")

	for i := len(companynameArray); i > 0; i-- {
		sem <- true 
		go func () { 
			defer func () { <- sem } () 
			_, err = stmt.Exec(companynameArray[i], companynumberArray[i], careofArray[i], poboxArray[i], addressline1Array[i], 
			addressline2Array[i], posttownArray[i], countyArray[i], countryArray[i], postcodeArray[i], 
			companycategoryArray[i], companystatusArray[i], countryoforiginArray[i], dissolutiondateArray[i], incorporatedateArray[i], 
			refdayArray[i], refmonthArray[i], a_nextduedateArray[i], a_lastmadeupdateArray[i], accountcategoryArray[i], 
			nextduedateArray[i], lastmadeupdateArray[i], mortchargesArray[i], mortoutstandingArray[i], mortpartsatisfiedArray[i], 
			mortsatisfiedArray[i], siccode1Array[i], siccode2Array[i], siccode3Array[i], siccode4Array[i], 
			genPartnerArray[i], limPartnerArray[i], uriArray[i], pn1_condate_Array[i], pn1_companyname_Array[i], 
			pn2_condate_Array[i], pn2_companyname_Array[i], pn3_condate_Array[i], pn3_companyname_Array[i], pn4_condate_Array[i], 
			pn4_companyname_Array[i], pn5_condate_Array[i], pn5_companyname_Array[i], pn6_condate_Array[i], pn6_companyname_Array[i], 
			pn7_condate_Array[i], pn7_companyname_Array[i], pn8_condate_Array[i], pn8_companyname_Array[i], pn9_condate_Array[i], 
			pn9_companyname_Array[i], pn10_condate_Array[i], pn10_companyname_Array[i], confstmtnextduedateArray[i], confstmtlastmadeupdateArray[i])
	
			checkErr(err, "Company Data Importation")
		}()
	}
}
