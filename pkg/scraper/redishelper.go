package scraper

import (
 	//"scraper"
 	"fmt"
 	"github.com/go-redis/redis"
)

var ini bool
var client *redis.Client

func RedisConnection() *redis.Client{
	fmt.Println("RedisConnection called ")

	if(ini == false){
		
		client = redis.NewClient(&redis.Options{
					Addr: "10.0.8.201:6379",
					Password: "", // no password set
					DB: 0,  // use default DB
				})
		_ , err := client.Ping().Result()
		if err != nil{
			panic(err)
		}
		ini = true
		fmt.Println("RedisConnection INITIALIZED")
	}

    return client
	// pong, err := client.Ping().Result()
	// fmt.Println(pong, err)

	// err = client.Set("key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := client.Get("key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// val2, err := client.Get("key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exists")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }
}
