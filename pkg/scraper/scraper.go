package scraper


import (
	"strings"
	"net/http"
	"io/ioutil"
	"fmt"
)

// type Scraper {
// 	Data_sources = [3]string{"seev", "drushim", "nisha"}
// }

func ProcessPost(post string){
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

func ScrapeBetween(data string, start string, end string) string{
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

func GetPageContent(url string) string{
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	return content
}

func getAllSkillNames() []string {
	return []string{"php", "mysql","api","rest", "json", "git","linux", "apache", "css3", "html5", "javascript"}
}
