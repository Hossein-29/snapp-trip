package databases

import (
	"example/snapp/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func ConnectToRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := Client.Ping().Result()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Successfully connected to Redis :)")
	}

	fmt.Println(pong)

	// err := client.Set("firstKey", "firstValue", 0).Err()
	// fmt.Println(err)
	// val, err := client.Get("firstKey").Result()
	// fmt.Println(val, err)

	//ctx := context.Background()

	// vaal1, err1 := client.Do("lpop", "Rule:THR-KRJ").Result()
	// if err1 != nil {
	// 	fmt.Println("hello", err1)
	// } else {
	// 	fmt.Println("ok", vaal1)
	// }
}

func CreateRouteSet(route string, ruleid int) {
	setName := "Rule:" + route
	val, err := Client.Do("SADD", setName, ruleid).Result()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("\nSeccussful redis : %d\n", val)
	}
}

func CreateAirlineSet(airline string, ruleid int) {
	setName := "RUle:" + airline
	val, err := Client.Do("SADD", setName, ruleid).Result()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("\nSeccussful redis : %d\n", val)
	}
}

func CreateAgencySet(agency string, ruleid int) {
	setName := "RUle:" + agency
	val, err := Client.Do("SADD", setName, ruleid).Result()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("\nSeccussful redis : %d\n", val)
	}
}

func CreateSupplierSet(supplier string, ruleid int) {
	setName := "RUle:" + supplier
	val, err := Client.Do("SADD", setName, ruleid).Result()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("\nSeccussful redis : %d\n", val)
	}
}

func MatchTicket(t models.Ticket, c *gin.Context) {
	routeSet := "Rule:" + t.Origin + "-" + t.Destination
	AirlineSet := "Rule:" + t.Airline
	AgencySet := "Rule:" + t.Agency
	SupplierSet := "Rule:" + t.Supplier
	val, err := Client.Do("SINTER", routeSet, AirlineSet, AgencySet, SupplierSet).Result()
	if err != nil {
		fmt.Println(err)
	}
	c.String(http.StatusOK, "Matched rules : %v", val)

}
