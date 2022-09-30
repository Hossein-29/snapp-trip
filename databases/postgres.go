package databases

import (
	"example/snapp/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Err error

func ConnectToPostgres() {
	dsn := "host=localhost user=postgres password=hb123456hb dbname=snapp port=5432 sslmode=disable"
	Db, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if Err != nil {
		fmt.Println(Err)
	} else {
		fmt.Println("Successfully connected to Postgres :)")
	}

	if created := GetValue("PostgresModelsCreated"); created == "false" {
		Db.AutoMigrate(&models.RulesTable{})
		Db.AutoMigrate(&models.RoutesTable{})
		Db.AutoMigrate(&models.AirlinesTable{})
		Db.AutoMigrate(&models.AgenciesTable{})
		Db.AutoMigrate(&models.SuppliersTable{})
		Db.AutoMigrate(&models.ValidCityTable{})
		Db.AutoMigrate(&models.ValidAirlineTable{})
		Db.AutoMigrate(&models.ValidAgencyTable{})
		Db.AutoMigrate(&models.ValidSupplierTable{})
		SetValue("PostgresModelsCreated", "true")
	}
}

func CreateRuleTable(t []models.Rule) {
	for i := range t {
		var RuleObj models.RulesTable
		RuleObj.AmountType = t[i].AmountType
		RuleObj.AmountValue = t[i].AmountValue
		Db.Model(&models.RulesTable{}).Select("AmountType", "AmountValue").Create(&RuleObj)
		CreateRuleHash(RuleObj)
		var RouteObj models.RoutesTable
		for _, j := range t[i].Routes {
			RouteObj.Route = j.Origin + "-" + j.Destination
			RouteObj.RuleId = RuleObj.Id
			CreateRouteSet(RouteObj.Route, RouteObj.RuleId)
			Db.Model(&models.RoutesTable{}).Select("Route", "RuleId").Create(&RouteObj)
		}
		var AirlineObj models.AirlinesTable
		for _, j := range t[i].Airlines {
			AirlineObj.Airline = j
			AirlineObj.RuleId = RuleObj.Id
			CreateAirlineSet(AirlineObj.Airline, AirlineObj.RuleId)
			Db.Model(&models.AirlinesTable{}).Select("Airline", "RuleId").Create(&AirlineObj)
		}
		var AgencyObj models.AgenciesTable
		for _, j := range t[i].Agencies {
			AgencyObj.Agency = j
			AgencyObj.RuleId = RuleObj.Id
			CreateAgencySet(AgencyObj.Agency, AirlineObj.RuleId)
			Db.Model(&models.AgenciesTable{}).Select("Agency", "RuleId").Create(&AgencyObj)
		}
		var SupplierObj models.SuppliersTable
		for _, j := range t[i].Suppliers {
			SupplierObj.Supplier = j
			SupplierObj.RuleId = RuleObj.Id
			CreateSupplierSet(SupplierObj.Supplier, SupplierObj.RuleId)
			Db.Model(&models.SuppliersTable{}).Select("Supplier", "RuleId").Create(&SupplierObj)
		}

	}
}

func CreateValidCityTable(city string) {
	var obj models.ValidCityTable
	obj.Name = city
	Db.Model(&models.ValidCityTable{}).Select("Name").Create(&obj)
}

func CreateValidAirlineTable(airline string) {
	var obj models.ValidAirlineTable
	obj.Name = airline
	Db.Model(&models.ValidAirlineTable{}).Select("Name").Create(&obj)
}

func CreateValidAgencyTable(agency string) {
	var obj models.ValidAgencyTable
	obj.Name = agency
	Db.Model(&models.ValidAgencyTable{}).Select("Name").Create(&obj)
}

func CreateValidSupplierTable(supplier string) {
	var obj models.ValidSupplierTable
	obj.Name = supplier
	Db.Model(&models.ValidSupplierTable{}).Select("Name").Create(&obj)
}

// This function could be used if we missed Redis to match tickets with rules
// Also for the best performance we should index name of routes, airlines,
// agencies and suppliers in their tables
/*
func MatchTicket(t models.Ticket) (report models.TicketResponse) {
	routeName1 := t.Origin + "-" + t.Destination
	routeName2 := t.Origin + "-"
	routeName3 := "-" + t.Destination
	routeName4 := "-"
	airlineName1 := t.Airline
	airlineName2 := ""
	agencyName1 := t.Agency
	agencyName2 := ""
	supplierName1 := t.Supplier
	supplierName2 := ""
	var matchedRules []models.RulesTable
	Db.Model(&models.RulesTable{}).
		Select("rules_tables.id, rules_tables.amount_type, rules_tables.amount_value").
		Joins("JOIN routes_tables ON rules_tables.id = routes_tables.rule_id AND (routes_tables.route = ? OR routes_tables.route = ? OR routes_tables.route = ? OR routes_tables.route = ?)", routeName1, routeName2, routeName3, routeName4).
		Joins("JOIN airlines_tables ON rules_tables.id = airlines_tables.rule_id AND (airlines_tables.airline = ? OR airlines_tables.airline = ?)", airlineName1, airlineName2).
		Joins("JOIN agencies_tables ON rules_tables.id = agencies_tables.rule_id AND (agencies_tables.agency = ? OR agencies_tables.agency = ?)", agencyName1, agencyName2).
		Joins("JOIN suppliers_tables ON rules_tables.id = suppliers_tables.rule_id AND (suppliers_tables.supplier = ? OR suppliers_tables.supplier = ?)", supplierName1, supplierName2).
		Find(&matchedRules)

	var basePrice float64 = t.BasePrice
	var bestMarkup float64 = 0
	var matchedRuleId int = -1

	for _, j := range matchedRules {
		ruleid := j.Id
		typeRule := j.AmountType
		valueRule := j.AmountValue
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
*/
