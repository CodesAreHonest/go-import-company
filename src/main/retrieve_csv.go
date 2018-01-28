package main

import (
	"fmt" 
	"strconv"
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"time"
)

func importCSVtoDB() { 
	
	start := time.Now()
	retrieveCSV()
	importDB()
	
	fmt.Printf("%.5fs seconds on import 4079779 rows from CSV to company_rawdata. \n", time.Since(start).Seconds())

}

func retrieveCSV() {
	

	csvFile, err := os.Open(COMPANY_FILE_DIRECTORY)
	checkErr(err, "Open CSV")

	defer csvFile.Close()

	// Create a new reader.
	reader := csv.NewReader(bufio.NewReader(csvFile))
	
	start := time.Now()

	for i := 0; i < ENTRIES; i++ {

		record, err := reader.Read()

		if i == 0 {
			continue
		}
		
		if i == 100000 { 
			fmt.Println("Imported 100000 rows", time.Since(start).Seconds()) 
		} else if i == 500000 { 
			fmt.Println("Imported 500000 rows", time.Since(start).Seconds())
		} else if i == 1000000 { 
			fmt.Println("Imported 1000000 rows", time.Since(start).Seconds())
		} else if i == 2000000 { 
			fmt.Println("Imported 2000000 rows", time.Since(start).Seconds())
		} else if i == 3000000 { 
			fmt.Println("Imported 3000000 rows", time.Since(start).Seconds())
		} else if i == 4000000 { 
			fmt.Println("Imported 4000000 rows", time.Since(start).Seconds())
		}

		// Stop at EOF.
		if err == io.EOF {
			break
		}
		
		int_mortchange, err := strconv.Atoi(record[22]) 
		checkErr(err, "convert mortchange value to integer") 
		
		int_mortoutstanding, err  := strconv.Atoi(record[23])
		checkErr(err, "convert mortoutstanding value to integer") 
		
		int_mortpartsatisfied, err := strconv.Atoi(record[24]) 
		checkErr(err, "convert mortpartsatisfied value to integer") 
		
		int_mortsatisfied, err  := strconv.Atoi(record[25])
		checkErr(err, "convert mortsatisfied value to integer") 
		
		int_genpartner, err := strconv.Atoi(record[30]) 
		checkErr(err, "convert genpartner value to integer")
		 
		int_limpartner, err := strconv.Atoi(record[31])
		checkErr(err, "convert limpartner value to integer")
		
		
		company := company_rawdata{	
			number: record[1],
			num_MortChanges: int_mortchange, 
			num_MortOutstanding: int_mortoutstanding, 
			num_MortPartSatisfied: int_mortpartsatisfied, 
			num_MortSatisfied: int_mortsatisfied,
			num_genPartner: int_genpartner,
			num_limPartner: int_limpartner,
			uri: record[32],
		}
		
		company.category.Scan(record[10])
		if len(company.category.String) == 0 {
			company.category.String = "Undefined"
		}
		
		company.status.Scan(record[11])
		if len(company.status.String) == 0 {
			company.status.String = "Undefined"
		}
		
		company.countryOfOrigin.Scan(record[12])
		if len(company.countryOfOrigin.String) < 2 {
			company.countryOfOrigin.String = "Undefined"
		}

		company.name.Scan(record[0])
		if len(company.name.String) == 0 {
			company.name.String = "Undefined"
		}

		company.careOf.Scan(record[2])
		if len(company.careOf.String) == 0 {
			company.careOf.String = "Undefined"
		}
		
		company.poBox.Scan(record[3])
		if len(company.poBox.String) == 0 {
			company.poBox.String = "Undefined"
		}
		
		company.addressLine1.Scan(record[4])
		if len(company.addressLine1.String) == 0 {
			company.addressLine1.String = "Undefined"
		}
		
		company.addressLine2.Scan(record[5])
		if len(company.addressLine2.String) == 0 {
			company.addressLine2.String = "Undefined"
		}
		
		company.postTown.Scan(record[6])
		if len(company.postTown.String) == 0 {
			company.postTown.String = "Undefined"
		}
		
		company.county.Scan(record[7])
		if len(company.county.String) == 0 {
			company.county.String = "Undefined"
		}
		
		company.country.Scan(record[8])
		if len(company.country.String) == 0 {
			company.country.String = "Undefined"
		}
		
		company.postcode.Scan(record[9])
		if len(company.postcode.String) == 0 {
			company.postcode.String = "Undefined"
		}
		
		company.dissolution_date.Scan(record[13])
		if len(company.dissolution_date.String) == 0 {
			company.dissolution_date.String = "01/01/3000"
		}
		
		company.incorporate_date.Scan(record[14])
		if len(company.dissolution_date.String) == 0 {
			company.dissolution_date.String = "01/01/3000"
		}
		
		company.accounting_refDay.Scan(record[15])
		company.accounting_refMonth.Scan(record[16])

		company.account_nextDueDate.Scan(record[17])
		if len(company.account_nextDueDate.String) == 0 {
			company.account_nextDueDate.String = "01/01/3000"
		}
		
		company.account_lastMadeUpdate.Scan(record[18])
		if len(company.account_lastMadeUpdate.String) == 0 {
			company.account_lastMadeUpdate.String = "01/01/3000"
		}
		
		company.account_category.Scan(record[19])
		if len(company.account_category.String) == 0 {
			company.account_category.String = "Undefined"
		}
		
		company.return_nextDueDate.Scan(record[20])
		if len(company.return_nextDueDate.String) == 0 {
			company.return_nextDueDate.String = "01/01/3000"
		}
		
		company.return_lastMadeUpdate.Scan(record[21])
		if len(company.return_lastMadeUpdate.String) == 0 {
			company.return_lastMadeUpdate.String = "01/01/3000"
		}
		
		company.siccode1.Scan(record[26])
		if len(company.siccode1.String) == 0 {
			company.siccode1.String = "Undefined"
		}
		
		company.siccode2.Scan(record[27])
		if len(company.siccode2.String) == 0 {
			company.siccode2.String = "Undefined"
		}
		
		company.siccode3.Scan(record[28])
		if len(company.siccode3.String) == 0 {
			company.siccode3.String = "Undefined"
		}
		
		company.siccode4.Scan(record[29])
		if len(company.siccode4.String) == 0 {
			company.siccode4.String = "Undefined"
		}
		
		company.pn1_condate.Scan(record[33])
		if len(company.pn1_condate.String) == 0 {
			company.pn1_condate.String = "01/01/3000"
		}
		
		company.pn1_companyname.Scan(record[34])
		if len(company.pn1_companyname.String) == 0 {
			company.pn1_companyname.String = "Undefined"
		}
		
		company.pn2_condate.Scan(record[35])
		if len(company.pn2_condate.String) == 0 {
			company.pn2_condate.String = "01/01/3000"
		
		company.pn2_companyname.Scan(record[36])
		if len(company.pn2_companyname.String) == 0 {
			company.pn2_companyname.String = "Undefined"
		}
		
		company.pn3_condate.Scan(record[37])
		if len(company.pn3_condate.String) == 0 {
			company.pn3_condate.String = "01/01/3000"
		}
		
		company.pn3_companyname.Scan(record[38])
		if len(company.pn3_companyname.String) == 0 {
			company.pn3_companyname.String = "Undefined"
		}
		
		company.pn4_condate.Scan(record[39])
		if len(company.pn4_condate.String) == 0 {
			company.pn4_condate.String = "01/01/3000"
		}
		
		company.pn4_companyname.Scan(record[40])
		if len(company.pn4_companyname.String) == 0 {
			company.pn4_companyname.String = "Undefined"
		}
		
		company.pn5_condate.Scan(record[41])
		if len(company.pn5_condate.String) == 0 {
			company.pn5_condate.String = "01/01/3000"
		}
		
		company.pn5_companyname.Scan(record[42])
		if len(company.pn5_companyname.String) == 0 {
			company.pn5_companyname.String = "Undefined"
		}
		
		company.pn6_condate.Scan(record[43])
		if len(company.pn6_condate.String) == 0 {
			company.pn6_condate.String = "01/01/3000"
		}
		
		company.pn6_companyname.Scan(record[44])
		if len(company.pn6_companyname.String) == 0 {
			company.pn6_companyname.String = "Undefined"
		}
		
		company.pn7_condate.Scan(record[45])
		if len(company.pn7_condate.String) == 0 {
			company.pn7_condate.String = "01/01/3000"
		}
		
		company.pn7_companyname.Scan(record[46])
		if len(company.pn7_companyname.String) == 0 {
			company.pn7_companyname.String = "Undefined"
		}
		
		company.pn8_condate.Scan(record[47])
		if len(company.pn8_condate.String) == 0 {
			company.pn8_condate.String = "01/01/3000"
		}
		
		company.pn8_companyname.Scan(record[48])
		if len(company.pn8_companyname.String) == 0 {
			company.pn8_companyname.String = "Undefined"
		}
		
		company.pn9_condate.Scan(record[49])
		if len(company.pn9_condate.String) == 0 {
			company.pn9_condate.String = "01/01/3000"
		}
		
		company.pn9_companyname.Scan(record[50])
		if len(company.pn9_companyname.String) == 0 {
			company.pn9_companyname.String = "Undefined"
		}
		
		company.pn10_condate.Scan(record[51])
		if len(company.pn10_condate.String) == 0 {
			company.pn10_condate.String = "01/01/3000"
		}
		
		company.pn10_companyname.Scan(record[52])
		if len(company.pn10_companyname.String) == 0 {
			company.pn10_companyname.String = "Undefined"
		}
		
		company.conf_stmtNextDueDate.Scan(record[53])
		if len(company.conf_stmtNextDueDate.String) == 0 {
			company.conf_stmtNextDueDate.String = "01/01/3000"
		}
		
		company.conf_stmtLastMadeUpdate.Scan(record[54])
		if len(company.conf_stmtLastMadeUpdate.String) == 0 {
			company.conf_stmtLastMadeUpdate.String = "01/01/3000"
		}
		
		companynameArray = append(companynameArray, company.name.String)
		companynumberArray = append(companynumberArray, company.number)
		careofArray = append(careofArray, company.careOf.String)
		poboxArray = append(poboxArray, company.poBox.String)
		addressline1Array = append(addressline1Array, company.addressLine1.String)
		
		addressline2Array = append(addressline2Array, company.addressLine2.String) 
		posttownArray = append(posttownArray, company.postTown.String) 
		countyArray = append(countyArray, company.county.String) 
		countryArray = append(countryArray, company.country.String) 
		postcodeArray = append(postcodeArray, company.postcode.String) 
		
		companycategoryArray = append(companycategoryArray, company.category.String) 
		companystatusArray = append(companystatusArray, company.status.String) 
		countryoforiginArray = append(countryoforiginArray, company.countryOfOrigin.String)
		dissolutiondateArray = append(dissolutiondateArray, company.dissolution_date.String) 
		incorporatedateArray = append(incorporatedateArray, company.incorporate_date.String) 
		
		refdayArray = append(refdayArray, company.accounting_refDay.Int64)
		refmonthArray = append(refmonthArray, company.accounting_refMonth.Int64) 
		a_nextduedateArray = append(a_nextduedateArray, company.account_nextDueDate.String) 
		a_lastmadeupdateArray = append(a_lastmadeupdateArray, company.account_lastMadeUpdate.String) 
		accountcategoryArray = append(accountcategoryArray, company.account_category.String) 
		
		nextduedateArray = append(nextduedateArray, company.return_nextDueDate.String) 
		lastmadeupdateArray = append(lastmadeupdateArray, company.return_lastMadeUpdate.String) 
		mortchargesArray = append(mortchargesArray, company.num_MortChanges) 
		mortoutstandingArray = append(mortoutstandingArray, company.num_MortOutstanding) 
		mortpartsatisfiedArray = append(mortpartsatisfiedArray, company.num_MortPartSatisfied)
		
		mortsatisfiedArray = append(mortsatisfiedArray, company.num_MortSatisfied)
		siccode1Array = append(siccode1Array, company.siccode1.String) 
		siccode2Array = append(siccode2Array, company.siccode2.String)
		siccode3Array = append(siccode3Array, company.siccode3.String)  
		siccode4Array = append(siccode4Array, company.siccode4.String) 
		 
		genPartnerArray = append(genPartnerArray, company.num_genPartner)
		limPartnerArray = append(limPartnerArray, company.num_limPartner) 
		uriArray = append(uriArray, company.uri) 
		pn1_condate_Array = append(pn1_condate_Array, company.pn1_condate.String)
		pn1_companyname_Array = append(pn1_companyname_Array, company.pn1_companyname.String)
		
		pn2_condate_Array = append(pn2_condate_Array, company.pn2_condate.String)
		pn2_companyname_Array = append(pn2_companyname_Array, company.pn2_companyname.String)
		pn3_condate_Array = append(pn3_condate_Array, company.pn3_condate.String)
		pn3_companyname_Array = append(pn3_companyname_Array, company.pn3_companyname.String)
		pn4_condate_Array = append(pn4_condate_Array, company.pn4_condate.String)
		
		pn4_companyname_Array = append(pn4_companyname_Array, company.pn4_companyname.String)
		pn5_condate_Array = append(pn5_condate_Array, company.pn5_condate.String)
		pn5_companyname_Array = append(pn5_companyname_Array, company.pn5_companyname.String)
		pn6_condate_Array = append(pn6_condate_Array, company.pn6_condate.String)
		pn6_companyname_Array = append(pn6_companyname_Array, company.pn6_companyname.String)
		
		pn7_condate_Array = append(pn7_condate_Array, company.pn7_condate.String)
		pn7_companyname_Array = append(pn7_companyname_Array, company.pn7_companyname.String)
		pn8_condate_Array = append(pn8_condate_Array, company.pn8_condate.String)
		pn8_companyname_Array = append(pn8_companyname_Array, company.pn8_companyname.String)
		pn9_condate_Array = append(pn9_condate_Array, company.pn9_condate.String)
		
		pn9_companyname_Array = append(pn9_companyname_Array, company.pn9_companyname.String)
		pn10_condate_Array = append(pn10_condate_Array, company.pn10_condate.String)
		pn10_companyname_Array = append(pn10_companyname_Array, company.pn10_companyname.String)
		confstmtnextduedateArray = append(confstmtnextduedateArray, company.conf_stmtNextDueDate.String)
		confstmtlastmadeupdateArray = append(confstmtlastmadeupdateArray, company.conf_stmtLastMadeUpdate.String)
		
			
	}
		
		
	}
	
	fmt.Println("company name ", len(companynameArray))
	fmt.Println("company number ", len(companynumberArray))		
	fmt.Println("company careof ", len(careofArray))		
	fmt.Println("company pobox ", len(poboxArray))	
	fmt.Println("company adrline1 ", len(addressline1Array))	
	
	fmt.Println("company adrline2 ", len(addressline2Array))
	fmt.Println("company posttwn ", len(posttownArray))		
	fmt.Println("company county ", len(countyArray))		
	fmt.Println("company country ", len(countryArray))	
	fmt.Println("company postcode ", len(postcodeArray))	
	
	fmt.Println("company category ", len(companycategoryArray))
	fmt.Println("company status ", len(companystatusArray))	
	fmt.Println("company con ", len(countryoforiginArray))		
	fmt.Println("company dis_date ", len(dissolutiondateArray))	
	fmt.Println("company incr_date ", len(incorporatedateArray))
	
	fmt.Println("company refday ", len(refdayArray))
	fmt.Println("company refmonth ", len(refmonthArray))	
	fmt.Println("company a_nextduedate ", len(a_nextduedateArray))		
	fmt.Println("company a_lastmadeupdate ", len(a_lastmadeupdateArray))	
	fmt.Println("company accountcategory ", len(accountcategoryArray))
	
	fmt.Println("company nextduedate ", len(nextduedateArray))
	fmt.Println("company lastmadeupdate ", len(lastmadeupdateArray))		
	fmt.Println("company mortcharge ", len(mortchargesArray))		
	fmt.Println("company mortoutstanding ", len(mortoutstandingArray))	
	fmt.Println("company mortpartsatisfied ", len(mortpartsatisfiedArray))
	
	fmt.Println("company mortsatisfied ", len(mortsatisfiedArray))
	fmt.Println("company siccode1 ", len(siccode1Array))
	fmt.Println("company siccode2 ", len(siccode2Array))
	fmt.Println("company siccode3 ", len(siccode3Array))
	fmt.Println("company siccode4 ", len(siccode4Array))
	
	fmt.Println("company genpartner ", len(genPartnerArray))
	fmt.Println("company limpartner ", len(limPartnerArray))		
	fmt.Println("company uri ", len(uriArray))		
	fmt.Println("company pn1_condate ", len(pn1_condate_Array))	
	fmt.Println("company pn1_companyname ", len(pn1_companyname_Array))
	
	fmt.Println("company pn2_condate ", len(pn2_condate_Array))	
	fmt.Println("company pn2_companyname ", len(pn2_companyname_Array))
	fmt.Println("company pn3_condate ", len(pn3_condate_Array))
	fmt.Println("company pn3_companyname ", len(pn3_companyname_Array))
	fmt.Println("company pn4_condate ", len(pn4_condate_Array))	
	fmt.Println("company pn4_companyname ", len(pn4_companyname_Array))
	fmt.Println("company pn5_condate ", len(pn5_condate_Array))		
	fmt.Println("company pn5_companyname ", len(pn5_companyname_Array))
	fmt.Println("company pn6_condate ", len(pn6_condate_Array))	
	fmt.Println("company pn6_companyname ", len(pn6_companyname_Array))
	fmt.Println("company pn7_condate ", len(pn7_condate_Array))		
	fmt.Println("company pn7_companyname ", len(pn7_companyname_Array))
	fmt.Println("company pn8_condate ", len(pn8_condate_Array))	
	fmt.Println("company pn8_companyname ", len(pn8_companyname_Array))
	fmt.Println("company pn9_condate ", len(pn9_condate_Array))		
	fmt.Println("company pn9_companyname ", len(pn9_companyname_Array))
	fmt.Println("company pn10_condate ", len(pn10_condate_Array))		
	fmt.Println("company pn10_companyname ", len(pn10_companyname_Array))
	fmt.Println("stmtnextduedate", len(confstmtnextduedateArray))
	fmt.Println("stmtlastmadeupdate", len(confstmtlastmadeupdateArray))	
}
