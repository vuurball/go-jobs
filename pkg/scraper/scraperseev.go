package scraper

import (
	"fmt"
//	 "os"
	"strings"
	"strconv"
//	"time"
	"sync"
)

type Scraperseev struct{

}

const(
	SITE_KEY = "seev"
	FIRST_URL = "http://www.seev.co.il/jobs/0-0-0/%D7%9B%D7%9C-%D7%94%D7%AA%D7%97%D7%95%D7%9E%D7%99%D7%9D/%D7%9B%D7%9C-%D7%94%D7%AA%D7%A4%D7%A7%D7%99%D7%93%D7%99%D7%9D/%D7%9B%D7%9C-%D7%94%D7%90%D7%96%D7%95%D7%A8%D7%99%D7%9D?page="
)

var wg sync.WaitGroup //process page posts as goroutine (will help only on the initial run)

/**
* implementing the SiteInterface
*/
func (this Scraperseev) Scrape(){
	nextPage := true
	page := 1
	url := FIRST_URL + strconv.Itoa(page)
	pagePostsHtml := []string{}
	redisClient := RedisConnection()
	
	for nextPage == true {
		fmt.Println("SCRAPING page" ,page)
			
	    pageFullHtml := GetPageContent(url)
		
		checkAndGetNextPage(pageFullHtml, &page, &url, &nextPage)

		wg.Add(1)
        go func() {
	        pagePostsHtml = strings.Split(pageFullHtml, "class=\"job row\"")
	        pagePostsHtml = append(pagePostsHtml[1:len(pagePostsHtml)]) //remove the first

	        for _, postContentHtml := range pagePostsHtml {
	        	if postKey := ""; postContentHtml != "" {
	        		postKey = ScrapeBetween(postContentHtml, "data-id=\"", "\">")

	 				sIsMember := redisClient.SIsMember(SITE_KEY, postKey)
	 				if sIsMember.Val() == false {
	        			ProcessPost(postContentHtml)
	        			redisClient.SAdd(SITE_KEY, postKey)
	        		}
	        	}
	        	//os.Exit(3)//stoping from looping during dev
	    	}
    		wg.Done()
    	}()
		//nextPage = false
	}	
	wg.Wait()
	

	fmt.Println("finished running", SITE_KEY)

}

/**
 * checks if next page exists, and creates the next url to go to, inc the current page, and nextPage pointer flags
 */
func checkAndGetNextPage(pageFullHtml string, page *int, url *string, nextPage *bool){

	if strings.Index(pageFullHtml, "<div class=\"job row") > -1{
	    *nextPage = true
	    *page++
	    *url = FIRST_URL + strconv.Itoa(*page)
	} else{
	    *nextPage = false // Setting nextPage to FALSE if there's no 'Next' link
	}
}
