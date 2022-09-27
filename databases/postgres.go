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

	if status := GetValue("PostgresModelsCreated"); status == "false" {
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
