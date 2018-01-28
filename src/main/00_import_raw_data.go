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

func importCSV() {
	initDB() 
	fmt.Println("Prepare to import data") 

	var sStmt string = "INSERT INTO company_rawdata VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55);"

	csvFile, err := os.Open(COMPANY_FILE_DIRECTORY)
	checkErr(err, "Open CSV")

	defer csvFile.Close()

	// Create a new reader.
	reader := csv.NewReader(bufio.NewReader(csvFile))
	
	start := time.Now()

	for i := 0; i <= ENTRIES; i++ {

		record, err := reader.Read()

		if i == 0 {
			continue
		}
		
		if i == 500 { 
			fmt.Println("Imported 500 rows", time.Since(start).Seconds()) 
		} else if i == 1000 { 
			fmt.Println("Imported 1000 rows", time.Since(start).Seconds())
		} else if i == 1500 { 
			fmt.Println("Imported 1500 rows", time.Since(start).Seconds())
		} else if i == 2000 { 
			fmt.Println("Imported 2000 rows", time.Since(start).Seconds())
		} else if i == 2500 { 
			fmt.Println("Imported 2500 rows", time.Since(start).Seconds())
		} else if i == 5000 { 
			fmt.Println("Imported 5000 rows", time.Since(start).Seconds())
		} else if i == 7500 { 
			fmt.Println("Will not display progress because too much row to process.", time.Since(start).Seconds())
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
		if len(company.countryOfOrigin.String) == 0 {
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
				
		_, err = db.Exec(sStmt, company.name.String, company.number, company.careOf.String, company.poBox.String, company.addressLine1.String, 
			company.addressLine2.String, company.postTown.String, company.county.String, company.country.String, company.postcode.String, 
			company.category.String, company.status.String, company.countryOfOrigin.String, company.dissolution_date.String, company.incorporate_date.String, 
			company.accounting_refDay.Int64, company.accounting_refMonth.Int64, company.account_nextDueDate.String, company.account_lastMadeUpdate.String, company.account_category.String, 
			company.return_nextDueDate.String, company.return_lastMadeUpdate.String, company.num_MortChanges, company.num_MortOutstanding, company.num_MortPartSatisfied, 
			company.num_MortSatisfied, company.siccode1.String, company.siccode2.String, company.siccode3.String, company.siccode4.String, 
			company.num_genPartner, company.num_limPartner, company.uri, company.pn1_condate.String, company.pn1_companyname.String, 
			company.pn2_condate.String, company.pn2_companyname.String, company.pn3_condate.String, company.pn3_companyname.String, company.pn4_condate.String, 
			company.pn4_companyname.String, company.pn5_condate.String, company.pn5_companyname.String, company.pn6_condate.String, company.pn6_companyname.String, 
			company.pn7_condate.String, company.pn7_companyname.String, company.pn8_condate.String, company.pn8_companyname.String, company.pn9_condate.String, 
			company.pn9_companyname.String, company.pn10_condate.String, company.pn10_companyname.String, company.conf_stmtNextDueDate.String, company.conf_stmtLastMadeUpdate.String)
		
		checkErr(err, "Company Data Importation")
	}
	
	
	}
	
	fmt.Printf("%.5fs seconds on import 4079779 rows from CSV to company_rawdata. \n", time.Since(start).Seconds())
	defer db.Close()
	
}


