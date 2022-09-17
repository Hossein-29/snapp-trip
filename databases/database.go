package databases

import (
	"example/snapp/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Err error

func ConnectToDatabase() {
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
	Db.AutoMigrate(&models.ValidCitiesTable{})
	Db.AutoMigrate(&models.ValidAirlinesTable{})
	Db.AutoMigrate(&models.ValidAgenciesTable{})
	Db.AutoMigrate(&models.ValidSuppliersTable{})

}
