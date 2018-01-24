package main

import (
	"database/sql"
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
	
	category 				string 
	status 					string 
	countryOfOrigin 		string 
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
