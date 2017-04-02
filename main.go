package main

import (
 	"scraper"
 	 . "fmt"
 	 "sync"
 	 "time"
)


var wg sync.WaitGroup

func main(){
	
	start := time.Now()
	dataSources := []string{"seev", "drushim", "nisha"}

	wg.Add(1)
	go func() {
		for _, dataSource := range dataSources{	
			s, ok:= scraper.GetScraperInstance("Scraper"+dataSource)
			if ok{
		    	s.Scrape()	
			}
		}
		wg.Done()
	}()

	wg.Wait()

	elapsed := time.Since(start)
    Println("time:",elapsed)
	//http.ListenAndServe(":8080", nil)    
}

 
