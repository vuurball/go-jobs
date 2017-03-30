package main

import (
	"fmt"
	"os"
	"net/http"
	"strings"
	 "strconv"
	 "io/ioutil"
	 "github.com/go-redis/redis"
)


func scrapeSeev(){

	key := "seev"
	firstUrl := "http://www.seev.co.il/jobs/0-0-0/%D7%9B%D7%9C-%D7%94%D7%AA%D7%97%D7%95%D7%9E%D7%99%D7%9D/%D7%9B%D7%9C-%D7%94%D7%AA%D7%A4%D7%A7%D7%99%D7%93%D7%99%D7%9D/%D7%9B%D7%9C-%D7%94%D7%90%D7%96%D7%95%D7%A8%D7%99%D7%9D?page="
	nextPage := true
	page := 1
	url := firstUrl + strconv.Itoa(page)
	pagePostsHtml := []string{}

	for nextPage == true {
		//fmt.Println("SCRAPING:" ,url)
		resp, _ := http.Get(url)
		bytes, _ := ioutil.ReadAll(resp.Body)
		pageFullHtml := string(bytes)

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
        		postKey = scrapeBetween(postContentHtml, "data-id=\"", "\">")
        		fmt.Println(postKey)
        		//if Redis::sismember(key, postKey) === 0 {
        			processPost(postContentHtml)
        		//}
        	}
        	os.Exit(3)
        }
		//fmt.Println("HTML:\n\n", string(bytes))
		nextPage = false
	}	

	fmt.Println("finished running", url, key)

}

func getAllSkillNames() []string {
	return []string{"php", "mysql","api","rest", "json", "git","linux", "apache", "css3", "html5", "javascript"}
}

func processPost(post string){
	//fmt.Println(post)
	fmt.Println("processed new post")
	// Redis::incr('postsCounter');
	skills := getAllSkillNames();
	var foundSkills []string
	post = strings.ToLower(post)

	for _, skill := range skills {
		if strings.Index(post, skill) > -1 {
			
			foundSkills = append(foundSkills, skill)
		}
	}

	fmt.Println("all skills found:",foundSkills)
	
	for len(foundSkills) > 1{
		 fmt.Println("node1:",foundSkills[0])
		 foundSkills = append(foundSkills[1:len(foundSkills)])
		 for _, skill2 := range foundSkills{
			fmt.Println(" connect with:",skill2)
		 }
	}
}

func scrapeBetween(data string, start string, end string) string{
 	startIndex := strings.Index(data, start)
    if startIndex == -1 {
        return ""
    }
    startIndex += len(start)
    endIndex := strings.Index(data, end)
    if endIndex == -1 {
        return ""
    }
    return data[startIndex:endIndex]
}

func main(){
    scrapeSeev()
	http.ListenAndServe(":8080", nil)    
}

func redisConnectionDraft(){
	client := redis.NewClient(&redis.Options{
		Addr: "10.0.8.201:6379",
		Password: "", // no password set
		DB: 0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
