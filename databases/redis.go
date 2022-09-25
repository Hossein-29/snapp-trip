package databases

import (
	"context"
	"example/snapp/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

var Client *redis.Client
var ctx context.Context

func ConnectToRedis() {
	ctx = context.Background()
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := Client.Ping(ctx).Result()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Successfully connected to Redis :)")
	}

	SetValue("ValidCityListCreated", "false")
	SetValue("ValidCityTableCreated", "false")
	SetValue("ValidAirlineListCreated", "false")
	SetValue("ValidAirlineTableCreated", "false")
	SetValue("ValidAgencyListCreated", "false")
	SetValue("ValidAgencyTableCreated", "false")
	SetValue("ValidSupplierListCreated", "false")
	SetValue("ValidSupplierTableCreated", "false")
}

func SetValue(key string, val interface{}) {
	_, err := Client.Set(ctx, key, val, 0).Result()
	if err != nil {
		fmt.Printf("SetValue: %s", err.Error())
	}
}

func GetValue(key string) string {
	res, err := Client.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("$$$GetValue: %s", err.Error())
	}
	return res
}

func CreateRuleHash(r models.RulesTable) {
	hashName := "Rule:" + strconv.Itoa(r.Id)
	_, err := Client.HSet(ctx, hashName, "amountType", r.AmountType, "amountValue", r.AmountValue).Result()
	if err != nil {
		fmt.Printf("CreateRuleHash: %s", err.Error())
	}
}

func CreateRouteSet(route string, ruleid int) {
	setName := "Rule:" + route
	_, err := Client.SAdd(ctx, setName, ruleid).Result()
	if err != nil {
		fmt.Printf("CreateRouteSet: %s", err.Error())
		return
	}
}

func CreateAirlineSet(airline string, ruleid int) {
	setName := "Rule:" + airline
	_, err := Client.SAdd(ctx, setName, ruleid).Result()
	if err != nil {
		fmt.Printf("CreateAirlineSet: %s", err.Error())
		return
	}
}

func CreateAgencySet(agency string, ruleid int) {
	setName := "Rule:" + agency
	_, err := Client.SAdd(ctx, setName, ruleid).Result()
	if err != nil {
		fmt.Printf("CreateAgencySet: %s", err.Error())
		return
	}
}

func CreateSupplierSet(supplier string, ruleid int) {
	setName := "Rule:" + supplier
	_, err := Client.SAdd(ctx, setName, ruleid).Result()
	if err != nil {
		fmt.Printf("CreateSupplierSet: %s", err.Error())
		return
	}
}

func CreateValidCityList(city string) {
	_, err := Client.RPush(ctx, "Valid:City", city).Result()
	if err != nil {
		fmt.Printf("CreateValidCityList: %s\n", err.Error())
		return
	}
}

func CreateValidAirlineList(airline string) {
	_, err := Client.RPush(ctx, "Valid:Airline", airline).Result()
	if err != nil {
		fmt.Printf("CreateValidAirlineList: %s\n", err.Error())
		return
	}
}

func CreateValidAgencyList(agency string) {
	_, err := Client.RPush(ctx, "Valid:Agency", agency).Result()
	if err != nil {
		fmt.Printf("CreateValidAgencyList: %s\n", err.Error())
		return
	}
}

func CreateValidSupplierList(supplier string) {
	_, err := Client.RPush(ctx, "Valid:Supplier", supplier).Result()
	if err != nil {
		fmt.Printf("CreateValidSupplierList: %s\n", err.Error())
		return
	}
}

func MatchTicket(t models.Ticket, c *gin.Context) (report models.TicketResponse) {
	routeSet := "Rule:" + t.Origin + "-" + t.Destination
	AirlineSet := "Rule:" + t.Airline
	AgencySet := "Rule:" + t.Agency
	SupplierSet := "Rule:" + t.Supplier
	matchedRules, err := Client.SInter(ctx, routeSet, AirlineSet, AgencySet, SupplierSet).Result()

	if err != nil {
		fmt.Printf("MatchTicket: %s", err.Error())
		return report
	}

	var basePrice float64 = t.BasePrice
	var bestMarkup float64 = 0
	var matchedRuleId int

	fmt.Print("\n\n length of matched rules : ", len(matchedRules), "\t")

	for i := range matchedRules {
		ruleid, _ := strconv.Atoi(matchedRules[i])
		hashName := "Rule:" + matchedRules[i]
		fmt.Printf("\n name of the hash : %s", hashName)
		typeRule, _ := Client.HGet(ctx, hashName, "amountType").Result()
		valueRule, _ := Client.HGet(ctx, hashName, "amountValue").Float64()
		fmt.Printf(" type : %s    value : %f", typeRule, valueRule)
		if typeRule == "FIXED" && valueRule > bestMarkup {
			fmt.Print("  hello1  ")
			matchedRuleId = ruleid
			bestMarkup = valueRule
		} else if typeRule == "PERCENTAGE" && (valueRule*basePrice/float64(100)) > bestMarkup {
			fmt.Print(" hello2  ")
			matchedRuleId = ruleid
			bestMarkup = (valueRule * basePrice / float64(100))
		} else {
			fmt.Print("  hello3  ")
		}
		fmt.Print(ruleid, "\t", bestMarkup, "\n\n")
	}

	report.RuleId = matchedRuleId
	report.Origin = t.Origin
	report.Airline = t.Airline
	report.Agency = t.Agency
	report.Supplier = t.Supplier
	report.BasePrice = t.BasePrice
	report.Markup = bestMarkup
	report.PayablePrice = basePrice + bestMarkup

	return report

}
