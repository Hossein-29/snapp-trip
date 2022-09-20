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
		fmt.Println("Successfully connected to database :)")
	}

	Db.AutoMigrate(&models.RulesTable{})
	Db.AutoMigrate(&models.RoutesTable{})
	Db.AutoMigrate(&models.AirlinesTable{})
	Db.AutoMigrate(&models.AgenciesTable{})
	Db.AutoMigrate(&models.SuppliersTable{})
	// Db.AutoMigrate(&models.ValidCitiesTable{})
	// Db.AutoMigrate(&models.ValidAirlinesTable{})
	// Db.AutoMigrate(&models.ValidAgenciesTable{})
	// Db.AutoMigrate(&models.ValidSuppliersTable{})
}

func CreateRuleTable(t []models.Rule) {
	for i := range t {
		var IdObj models.RulesTable
		Db.Model(&models.RulesTable{}).Create(&IdObj)
		var RouteObj models.RoutesTable
		for _, j := range t[i].Routes {
			RouteObj.Route = j.Origin + "-" + j.Destination
			RouteObj.RuleId = IdObj.Id
			Db.Model(&models.RoutesTable{}).Select("Route", "RuleId").Create(&RouteObj)
		}
		var AirlineObj models.AirlinesTable
		for _, j := range t[i].Airlines {
			AirlineObj.Airline = j
			Db.Model(&models.AirlinesTable{}).Select("Airline").Create(&AirlineObj)
		}
		var AgencyObj models.AgenciesTable
		for _, j := range t[i].Agencies {
			AgencyObj.Agency = j
			Db.Model(&models.AgenciesTable{}).Select("Agency").Create(&AgencyObj)
		}
		var SupplierObj models.SuppliersTable
		for _, j := range t[i].Suppliers {
			SupplierObj.Supplier = j
			Db.Model(&models.SuppliersTable{}).Select("Supplier").Create(&SupplierObj)
		}
	}
}
