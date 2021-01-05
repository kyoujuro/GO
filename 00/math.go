package main

import (
	"fmt"
	"time
	)
	
func main() {
	
	println(12 + 3)
    
    	println(3 - 6)
    
    	println(5 * 9)

	println(9 % 4)

    
    	t := time.Now()
    	fmt.Println(t) 

   
   	fmt.Println(t)              
	fmt.Println(t.Year())       
  	fmt.Println(t.YearDay())    
    	fmt.Println(t.Month())      
    	fmt.Println(t.Weekday())    
	fmt.Println(t.Day())        
    	fmt.Println(t.Hour())       
    	fmt.Println(t.Minute())     
    	fmt.Println(t.Second())     
    	fmt.Println(t.Nanosecond()) 
    	fmt.Println(t.Zone())      
	}
    
}
