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
	var cnt int64

	// city
	if cnt, _ = Client.Exists(ctx, "ValidCityListCreated").Result(); cnt == 0 {
		SetValue("ValidCityListCreated", "false")
	}
	if cnt, _ = Client.Exists(ctx, "ValidCityTableCreated").Result(); cnt == 0 {
		SetValue("ValidCityTableCreated", "false")
	}

	// airline
	if cnt, _ = Client.Exists(ctx, "ValidAirlineListCreated").Result(); cnt == 0 {
		SetValue("ValidAirlineListCreated", "false")
	}
	if cnt, _ = Client.Exists(ctx, "ValidAirlineTableCreated").Result(); cnt == 0 {
		SetValue("ValidAirlineTableCreated", "false")
	}

	// agency
	if cnt, _ = Client.Exists(ctx, "ValidAgencyListCreated").Result(); cnt == 0 {
		SetValue("ValidAgencyListCreated", "false")
	}
	if cnt, _ = Client.Exists(ctx, "ValidAgencyTableCreated").Result(); cnt == 0 {
		SetValue("ValidAgencyTableCreated", "false")
	}

	// supplier
	if cnt, _ = Client.Exists(ctx, "ValidSupplierListCreated").Result(); cnt == 0 {
		SetValue("ValidSupplierListCreated", "false")
	}
	if cnt, _ = Client.Exists(ctx, "ValidSupplierTableCreated").Result(); cnt == 0 {
		SetValue("ValidSupplierTableCreated", "false")
	}

	// postgres models
	if cnt, _ = Client.Exists(ctx, "PostgresModelsCreated").Result(); cnt == 0 {
		SetValue("PostgresModelsCreated", "false")
	}

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
		fmt.Printf("GetValue: %s", err.Error())
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
	setName := "Rule:Route:" + route
	_, err := Client.SAdd(ctx, setName, ruleid).Result()
	if err != nil {
		fmt.Printf("CreateRouteSet: %s", err.Error())
		return
	}
}

func CreateAirlineSet(airline string, ruleid int) {
	setName := "Rule:Airline:" + airline
	_, err := Client.SAdd(ctx, setName, ruleid).Result()
	if err != nil {
		fmt.Printf("CreateAirlineSet: %s", err.Error())
		return
	}
}

func CreateAgencySet(agency string, ruleid int) {
	setName := "Rule:Agency:" + agency
	_, err := Client.SAdd(ctx, setName, ruleid).Result()
	if err != nil {
		fmt.Printf("CreateAgencySet: %s", err.Error())
		return
	}
}

func CreateSupplierSet(supplier string, ruleid int) {
	setName := "Rule:Supplier:" + supplier
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
	routeSet1 := "Rule:Route:" + t.Origin + "-" + t.Destination
	routeSet2 := "Rule:Route:" + t.Origin + "-"
	routeSet3 := "Rule:Route:" + "-" + t.Destination
	routeSet4 := "Rule:Route:" + "-"
	airlineSet1 := "Rule:Airline:" + t.Airline
	airlineSet2 := "Rule:Airline:"
	agencySet1 := "Rule:Agency:" + t.Agency
	agencySet2 := "Rule:Agency:"
	supplierSet1 := "Rule:Supplier:" + t.Supplier
	supplierSet2 := "Rule:Supplier:"

	Client.SUnionStore(ctx, "TempRoute", routeSet1, routeSet2, routeSet3, routeSet4)
	Client.SUnionStore(ctx, "TempAirline", airlineSet1, airlineSet2)
	Client.SUnionStore(ctx, "TempAgency", agencySet1, agencySet2)
	Client.SUnionStore(ctx, "TempSupplier", supplierSet1, supplierSet2)

	matchedRules, err := Client.SInter(ctx, "TempRoute", "TempAirline", "TempAgency", "TempSupplier").Result()
	if err != nil {
		fmt.Printf("MatchTicket: %s", err.Error())
		return report
	}

	var basePrice float64 = t.BasePrice
	var bestMarkup float64 = 0
	var matchedRuleId int = -1

	for i := range matchedRules {
		ruleid, _ := strconv.Atoi(matchedRules[i])
		hashName := "Rule:" + matchedRules[i]
		typeRule, _ := Client.HGet(ctx, hashName, "amountType").Result()
		valueRule, _ := Client.HGet(ctx, hashName, "amountValue").Float64()
		if typeRule == "FIXED" && valueRule > bestMarkup {
			matchedRuleId = ruleid
			bestMarkup = valueRule
		} else if typeRule == "PERCENTAGE" && (valueRule*basePrice/float64(100)) > bestMarkup {
			matchedRuleId = ruleid
			bestMarkup = (valueRule * basePrice / float64(100))
		}
	}

	report.RuleId = matchedRuleId
	report.Origin = t.Origin
	report.Destination = t.Destination
	report.Airline = t.Airline
	report.Agency = t.Agency
	report.Supplier = t.Supplier
	report.BasePrice = basePrice
	report.Markup = bestMarkup
	report.PayablePrice = basePrice + bestMarkup

	return report

}
