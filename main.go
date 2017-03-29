package main

import (
	"fmt"
	"net/http"
	"strings"
	 "strconv"
	 "io/ioutil"
)


func scrapeSeev(){

	key := "seev"
	firstUrl := "http://www.seev.co.il/jobs/0-0-0/%D7%9B%D7%9C-%D7%94%D7%AA%D7%97%D7%95%D7%9E%D7%99%D7%9D/%D7%9B%D7%9C-%D7%94%D7%AA%D7%A4%D7%A7%D7%99%D7%93%D7%99%D7%9D/%D7%9B%D7%9C-%D7%94%D7%90%D7%96%D7%95%D7%A8%D7%99%D7%9D?page="
	nextPage := true
	page := 1
	url := firstUrl + strconv.Itoa(page)

	for nextPage == true{
		fmt.Println("SCRAPING:" ,url)
		resp, _ := http.Get(url)
		bytes, _ := ioutil.ReadAll(resp.Body)
		pageHtml :=string(bytes)
		//fmt.Println("HTML:\n\n", string(bytes))
		nextPage = false
	}	

	fmt.Println("loaded", url, key)
	skills := [7]string{"php", "mysql", "linux", "apache", "css3", "html5", "javascript"}
	//fmt.Println(skills)

	post := "<div class=\"jobFields\"><spancss3 php mysql Linux class=\"sendCV\"><a action=\"https://www.drushim.co.il/sendcv.a"
	//fmt.Println(post)

	var foundSkills []string
	fmt.Println(foundSkills)
	
	//fmt.Println(strings.ToLower(post))
	post = strings.ToLower(post)

	for _, skill := range skills {
		if strings.Index(post, skill) > -1 {
			//fmt.Println("skill", skill, "found")
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

func main(){
	scrapeSeev()
	http.ListenAndServe(":8080", nil)    
}
