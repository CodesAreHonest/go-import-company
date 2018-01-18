package main

import (
	//	"fmt"
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

func openCSV() {

	var sStmt string = "insert into company_test values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55)"

	stmt, err := db.Prepare(sStmt)
	checkErr(err, "Prepare COmpany Stmt")

	csvFile, err := os.Open(COMPANY_FILE_DIRECTORY)
	checkErr(err, "Open CSV")

	defer csvFile.Close()

	// Create a new reader.
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for i := 0; i <= 5; i++ {

		record, err := reader.Read()

		if i == 0 {
			continue
		}

		// Stop at EOF.
		if err == io.EOF {
			break
		}

		company := company_rawdata{
			number: record[1],
			category: record[10],
			status: record[11], 
			countryOfOrigin: record[12], 
			num_MortChanges: record[22], 
			num_MortOutstanding: record[23], 
			num_MortPartSatisfied: record[24], 
			num_MortSatisfied: record[25],
			num_genPartner: record[30],
			num_limPartner: record[31],
			uri: record[32],
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
			company.dissolution_date.String = "3000-01-01"
		}
		
		company.incorporate_date.Scan(record[14])
		if len(company.dissolution_date.String) == 0 {
			company.dissolution_date.String = "3000-01-01"
		}
		
		company.accounting_refDay.Scan(record[15])
		company.accounting_refMonth.Scan(record[16])

		company.account_nextDueDate.Scan(record[17])
		if len(company.account_nextDueDate.String) == 0 {
			company.account_nextDueDate.String = "3000-01-01"
		}
		
		company.account_lastMadeUpdate.Scan(record[18])
		if len(company.account_lastMadeUpdate.String) == 0 {
			company.account_lastMadeUpdate.String = "3000-01-01"
		}
		
		company.account_category.Scan(record[19])
		if len(company.account_category.String) == 0 {
			company.account_category.String = "Undefined"
		}
		
		company.return_nextDueDate.Scan(record[20])
		if len(company.return_nextDueDate.String) == 0 {
			company.return_nextDueDate.String = "3000-01-01"
		}
		
		company.return_lastMadeUpdate.Scan(record[21])
		if len(company.return_lastMadeUpdate.String) == 0 {
			company.return_lastMadeUpdate.String = "3000-01-01"
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
			company.pn1_condate.String = "3000-01-01"
		}
		
		company.pn1_companydate.Scan(record[34])
		if len(company.pn1_companydate.String) == 0 {
			company.pn1_companydate.String = "Undefined"
		}
		
		company.pn2_condate.Scan(record[35])
		if len(company.pn2_condate.String) == 0 {
			company.pn2_condate.String = "3000-01-01"
		}
		
		company.pn2_companyname.Scan(record[36])
		if len(company.pn2_companyname.String) == 0 {
			company.pn2_companyname.String = "Undefined"
		}
		
		company.pn3_condate.Scan(record[37])
		if len(company.pn3_condate.String) == 0 {
			company.pn3_condate.String = "3000-01-01"
		}
		
		company.pn3_companyname.Scan(record[38])
		if len(company.pn3_companyname.String) == 0 {
			company.pn3_companyname.String = "Undefined"
		}
		
		company.pn4_condate.Scan(record[39])
		if len(company.pn4_condate.String) == 0 {
			company.pn4_condate.String = "3000-01-01"
		}
		
		company.pn4_companyname.Scan(record[40])
		if len(company.pn4_companyname.String) == 0 {
			company.pn4_companyname.String = "Undefined"
		}
		
		company.pn5_condate.Scan(record[37])
		if len(company.pn5_condate.String) == 0 {
			company.pn5_condate.String = "3000-01-01"
		}
		
		company.pn5_companyname.Scan(record[38])
		if len(company.pn5_companyname.String) == 0 {
			company.pn5_companyname.String = "Undefined"
		}
		
		company.pn6_condate.Scan(record[39])
		if len(company.pn6_condate.String) == 0 {
			company.pn6_condate.String = "3000-01-01"
		}
		
		company.pn6_companyname.Scan(record[40])
		if len(company.pn6_companyname.String) == 0 {
			company.pn6_companyname.String = "Undefined"
		}
		
		company.pn7_condate.Scan(record[37])
		if len(company.pn7_condate.String) == 0 {
			company.pn7_condate.String = "3000-01-01"
		}
		
		company.pn7_companyname.Scan(record[38])
		if len(company.pn7_companyname.String) == 0 {
			company.pn7_companyname.String = "Undefined"
		}
		
		company.pn8_condate.Scan(record[39])
		if len(company.pn8_condate.String) == 0 {
			company.pn8_condate.String = "3000-01-01"
		}
		
		company.pn8_companyname.Scan(record[40])
		if len(company.pn8_companyname.String) == 0 {
			company.pn8_companyname.String = "Undefined"
		}
		
		company.pn9_condate.Scan(record[37])
		if len(company.pn9_condate.String) == 0 {
			company.pn9_condate.String = "3000-01-01"
		}
		
		company.pn9_companyname.Scan(record[38])
		if len(company.pn9_companyname.String) == 0 {
			company.pn9_companyname.String = "Undefined"
		}
		
		company.pn10_condate.Scan(record[39])
		if len(company.pn10_condate.String) == 0 {
			company.pn10_condate.String = "3000-01-01"
		}
		
		company.pn10_companyname.Scan(record[40])
		if len(company.pn10_companyname.String) == 0 {
			company.pn10_companyname.String = "Undefined"
		}
		
		company.conf_stmtNextDueDate.Scan(record[39])
		if len(company.conf_stmtNextDueDate.String) == 0 {
			company.conf_stmtNextDueDate.String = "3000-01-01"
		}
		
		company.conf_stmtLastMadeUpdate.Scan(record[40])
		if len(company.conf_stmtLastMadeUpdate.String) == 0 {
			company.conf_stmtLastMadeUpdate.String = "3000-01-01"
		}

		stmt.Exec(company.companyname.String, company.companynumber, company.careof.String)
		checkErr(err, "Company Data Importation")
	}
}
