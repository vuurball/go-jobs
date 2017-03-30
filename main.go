package main

import (
 	"scraper"
 	//"github.com/go-redis/redis"
)


func main(){
	s := scraper.Scraperseev{}
    s.Scrape()
	//http.ListenAndServe(":8080", nil)    
}

// func redisConnectionDraft(){
// 	client := redis.NewClient(&redis.Options{
// 		Addr: "10.0.8.201:6379",
// 		Password: "", // no password set
// 		DB: 0,  // use default DB
// 	})

// 	pong, err := client.Ping().Result()
// 	fmt.Println(pong, err)

// 	err = client.Set("key", "value", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, err := client.Get("key").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := client.Get("key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exists")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// }
