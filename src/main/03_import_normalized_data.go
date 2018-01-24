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
	insert_distinct_partnership() 
	
	fmt.Printf("%.5fs seconds on import partnership. \n", time.Since(start).Seconds())
}