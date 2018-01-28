package main

import (
	"database/sql"
)

var ( 
	uriArray [] string 
	
	genPartnerArray [] int 
	limPartnerArray [] int 
	
	mortchargesArray [] int 
	mortoutstandingArray [] int 
	mortpartsatisfiedArray [] int 
	mortsatisfiedArray [] int 
	
	siccode1Array [] string
	siccode2Array [] string
	siccode3Array [] string
	siccode4Array [] string
	
	nextduedateArray [] string 
	lastmadeupdateArray [] string 
	
	categoryArray[] string
	accountcategoryArray[] string
	
	refdayArray [] int64 
	refmonthArray [] int64 
	a_nextduedateArray[] string 
	a_lastmadeupdateArray [] string 
	categoryIDArray [] int 
	
	confstmtnextduedateArray[] string 
	confstmtlastmadeupdateArray[] string 
	
	careofArray [] string 
	poboxArray [] string 
	addressline1Array [] string 
	addressline2Array [] string 
	posttownArray [] string 
	countyArray [] string 
	countryArray [] string 
	postcodeArray [] string 
	
	countryoforiginArray [] string
	companystatusArray [] string  
	companycategoryArray [] string 
	
	dissolutiondateArray [] string 
	incorporatedateArray [] string 
	
	companynameArray [] string 
	companynumberArray [] string 
	
	com_address_idArray [] int 
	com_category_idArray [] int 
	com_status_idArray [] int 
	com_countryoforigin_idArray[] int 
	
	pn1_condate_Array [] string 
	pn1_companyname_Array [] string 
	pn2_condate_Array [] string 
	pn2_companyname_Array [] string 
	pn3_condate_Array [] string 
	pn3_companyname_Array [] string 
	pn4_condate_Array [] string 
	pn4_companyname_Array [] string 
	pn5_condate_Array [] string 
	pn5_companyname_Array [] string 
	pn6_condate_Array [] string 
	pn6_companyname_Array [] string 
	pn7_condate_Array [] string 
	pn7_companyname_Array [] string 
	pn8_condate_Array [] string 
	pn8_companyname_Array [] string 
	pn9_condate_Array [] string 
	pn9_companyname_Array [] string 
	pn10_condate_Array [] string 
	pn10_companyname_Array [] string 
	
	// for import company table 
	com_detail_idArray [] int
	com_acc_idArray [] int
	com_return_idArray [] int 
	com_mort_idArray [] int 
	com_sic_idArray [] int
	com_partnership_idArray [] int 
	com_uri_idArray [] int
	com_previousname_idArray [] int  
	com_conf_stmt_idArray [] int

)



type company_rawdata struct { 
	
	name 					sql.NullString
	number 					string 
	careOf 					sql.NullString 
	poBox 					sql.NullString  
	addressLine1 			sql.NullString // 5
	
	addressLine2 			sql.NullString  
	postTown 				sql.NullString  
	county 					sql.NullString  
	country 				sql.NullString  
	postcode 				sql.NullString // 10
	
	category 				sql.NullString 
	status 					sql.NullString 
	countryOfOrigin 		sql.NullString 
	dissolution_date 		sql.NullString
	incorporate_date 		sql.NullString // 15
	 
	accounting_refDay 		sql.NullInt64
	accounting_refMonth 	sql.NullInt64
	account_nextDueDate 	sql.NullString
	account_lastMadeUpdate 	sql.NullString 
	account_category 		sql.NullString // 20
	
	return_nextDueDate 		sql.NullString
	return_lastMadeUpdate 	sql.NullString 
	num_MortChanges 		int 
	num_MortOutstanding 	int
	num_MortPartSatisfied 	int // 25
	
	num_MortSatisfied 		int
	siccode1 				sql.NullString 
	siccode2 				sql.NullString
	siccode3 				sql.NullString
	siccode4 				sql.NullString // 30
	
	num_genPartner 			int
	num_limPartner 			int
	uri 					string 
	pn1_condate 			sql.NullString 
	pn1_companyname 		sql.NullString // 35
	
	pn2_condate 			sql.NullString 
	pn2_companyname 		sql.NullString
	pn3_condate 			sql.NullString 
	pn3_companyname 		sql.NullString
	pn4_condate 			sql.NullString // 40
	
	pn4_companyname 		sql.NullString
	pn5_condate 			sql.NullString 
	pn5_companyname 		sql.NullString
	pn6_condate 			sql.NullString 
	pn6_companyname 		sql.NullString // 45
	
	pn7_condate 			sql.NullString 
	pn7_companyname 		sql.NullString
	pn8_condate 			sql.NullString 
	pn8_companyname 		sql.NullString
	pn9_condate 			sql.NullString // 50
	
	pn9_companyname 		sql.NullString
	pn10_condate 			sql.NullString 
	pn10_companyname 		sql.NullString
	conf_stmtNextDueDate 	sql.NullString 
	conf_stmtLastMadeUpdate sql.NullString // 55
	
}



//type NullInt64 sql.NullInt64
//
//// Scan implements the Scanner interface for NullInt64
//func (ni *NullInt64) Scan(value interface{}) error {
//	var i sql.NullInt64
//	if err := i.Scan(value); err != nil {
//		return err
//	}
//
//	// if nil then make Valid false
//	if reflect.TypeOf(value) == nil {
//		*ni = NullInt64{i.Int64, false}
//	} else {
//		*ni = NullInt64{i.Int64, true}
//	}
//	return nil
//}
//
//// NullBool is an alias for sql.NullBool data type
//type NullBool sql.NullBool
//
//// Scan implements the Scanner interface for NullBool
//func (nb *NullBool) Scan(value interface{}) error {
//	var b sql.NullBool
//	if err := b.Scan(value); err != nil {
//		return err
//	}
//
//	// if nil then make Valid false
//	if reflect.TypeOf(value) == nil {
//		*nb = NullBool{b.Bool, false}
//	} else {
//		*nb = NullBool{b.Bool, true}
//	}
//
//	return nil
//}
//
//// NullFloat64 is an alias for sql.NullFloat64 data type
//type NullFloat64 sql.NullFloat64
//
//// Scan implements the Scanner interface for NullFloat64
//func (nf *NullFloat64) Scan(value interface{}) error {
//	var f sql.NullFloat64
//	if err := f.Scan(value); err != nil {
//		return err
//	}
//
//	// if nil then make Valid false
//	if reflect.TypeOf(value) == nil {
//		*nf = NullFloat64{f.Float64, false}
//	} else {
//		*nf = NullFloat64{f.Float64, true}
//	}
//
//	return nil
//}
//
//// NullString is an alias for sql.NullString data type
//type NullString sql.NullString
//
//// Scan implements the Scanner interface for NullString
//func (ns *NullString) Scan(value interface{}) error {
//	var s sql.NullString
//	if err := s.Scan(value); err != nil {
//		return err
//	}
//
//	// if nil then make Valid false
//	if reflect.TypeOf(value) == nil {
//		*ns = NullString{s.String, false}
//	} else {
//		*ns = NullString{s.String, true}
//	}
//
//	return nil
//}
