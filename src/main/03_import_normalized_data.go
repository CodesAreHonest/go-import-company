package main

import (
	"time"
	"fmt"
)

func import_uri() { 
	
	start := time.Now()
	retrieve_distinct_uri()
	insert_distinct_uri() 
	
	fmt.Printf("%.5fs seconds on import uri. \n", time.Since(start).Seconds())
}

func import_partnership() { 
	
	start := time.Now()
	retrieve_distinct_partnership()
	insert_distinct_partnership() 
	
	fmt.Printf("%.5fs seconds on import partnership. \n", time.Since(start).Seconds())
}

func import_mortgages() { 
	
	start := time.Now()
	retrieve_distinct_mortgages()
	insert_distinct_mortgages() 
	
	fmt.Printf("%.5fs seconds on import mortgages. \n", time.Since(start).Seconds())
}

func import_returns() { 
	
	start := time.Now()
	retrieve_distinct_returns()
	insert_distinct_returns() 
	
	fmt.Printf("%.5fs seconds on import returns. \n", time.Since(start).Seconds())
}

func import_account_category() { 
	
	start := time.Now()
	retrieve_distinct_account_category()
	insert_distinct_account_category() 
	
	fmt.Printf("%.5fs seconds on import account category. \n", time.Since(start).Seconds())
}

func import_account() { 
	
	start := time.Now()
	retrieve_distinct_account()
	insert_distinct_account() 
	
	fmt.Printf("%.5fs seconds on import account. \n", time.Since(start).Seconds())
}

func import_conference_statement() { 
	
	start := time.Now()
	retrieve_distinct_conf_statement()
	insert_distinct_conf_statement() 
	
	fmt.Printf("%.5fs seconds on import conference statement. \n", time.Since(start).Seconds())
}

func import_countryoforigin() { 
	
	start := time.Now()
	retrieve_distinct_countryoforigin()
	insert_distinct_countryoforigin() 
	
	fmt.Printf("%.5fs seconds on import countryoforigin statement. \n", time.Since(start).Seconds())
}

func import_companystatus() { 
	
	start := time.Now()
	retrieve_distinct_companystatus()
	insert_distinct_companystatus() 
	
	fmt.Printf("%.5fs seconds on import companystatus statement. \n", time.Since(start).Seconds())
}

func import_companycategory() { 
	
	start := time.Now()
	retrieve_distinct_companycategory()
	insert_distinct_companycategory() 
	
	fmt.Printf("%.5fs seconds on import company category statement. \n", time.Since(start).Seconds())
}

func import_companydetail() { 
	
	start := time.Now()
	retrieve_distinct_companydetail()
	insert_distinct_companydetail() 
	
	fmt.Printf("%.5fs seconds on import companydetail statement. \n", time.Since(start).Seconds())
}

func import_companysiccode() { 
	
	start := time.Now()
	retrieve_distinct_companysiccode()
	insert_distinct_companysiccode() 
	
	fmt.Printf("%.5fs seconds on import companysiccode statement. \n", time.Since(start).Seconds())
}

func import_previousname() { 
	
	start := time.Now()
	retrieve_distinct_previousname()
	insert_distinct_previousname() 
	
	fmt.Printf("%.5fs seconds on import company previousname statement. \n", time.Since(start).Seconds())
}