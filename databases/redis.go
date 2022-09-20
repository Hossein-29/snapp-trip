package databases

import (
	"fmt"

	"github.com/go-redis/redis"
)

func ConnectToRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//pong, err := client.Ping().Result()

	err := client.Set("firstKey", "firstValue", 0).Err()
	fmt.Println(err)
	val, err := client.Get("firstKey").Result()
	fmt.Println(val, err)

	//ctx := context.Background()

	// vaal1, err1 := client.Do("lpop", "Rule:THR-KRJ").Result()
	// if err1 != nil {
	// 	fmt.Println("hello", err1)
	// } else {
	// 	fmt.Println("ok", vaal1)
	// }
}
