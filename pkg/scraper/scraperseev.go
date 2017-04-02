package scraper

import (
	"fmt"
//	 "os"
	"strings"
	"strconv"
)

type Scraperseev struct{

}

/**
* implementing the SiteInterface
*/
func (this Scraperseev) Scrape(){

	key := "seev"
	firstUrl := "http://www.seev.co.il/jobs/0-0-0/%D7%9B%D7%9C-%D7%94%D7%AA%D7%97%D7%95%D7%9E%D7%99%D7%9D/%D7%9B%D7%9C-%D7%94%D7%AA%D7%A4%D7%A7%D7%99%D7%93%D7%99%D7%9D/%D7%9B%D7%9C-%D7%94%D7%90%D7%96%D7%95%D7%A8%D7%99%D7%9D?page="
	nextPage := true
	page := 1
	url := firstUrl + strconv.Itoa(page)
	pagePostsHtml := []string{}

	redisClient := RedisConnection()

	for nextPage == true {
		fmt.Println("SCRAPING:" ,url)
		pageFullHtml := GetPageContent(url)

		if strings.Index(pageFullHtml, "<div class=\"job row") > -1{
            nextPage = true
            page++
            url = firstUrl + strconv.Itoa(page)
            //fmt.Println("next page found")
        } else{
        	//fmt.Println("next page not found")
            nextPage = false // Setting nextPage to FALSE if there's no 'Next' link
        }

        pagePostsHtml = strings.Split(pageFullHtml, "class=\"job row\"")
        pagePostsHtml = append(pagePostsHtml[1:len(pagePostsHtml)]) //remove the first
        //fmt.Println(len(pagePostsHtml))

        for _, postContentHtml := range pagePostsHtml {
        	if postKey := ""; postContentHtml != "" {
        		postKey = ScrapeBetween(postContentHtml, "data-id=\"", "\">")

 				sIsMember := redisClient.SIsMember(key, postKey)
 				if sIsMember.Val() == false {
        			ProcessPost(postContentHtml)
        			redisClient.SAdd(key, postKey)
        		}
        	}
        //	os.Exit(3)//stoping from looping during dev
        }
		
		//nextPage = false
	}	

	fmt.Println("finished running", url, key)

}
