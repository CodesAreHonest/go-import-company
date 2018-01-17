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