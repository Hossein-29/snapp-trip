package databases

import (
	"example/snapp/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Err error

func ConnectToDatabase() {
	dsn := "host=localhost user=postgres password=hb123456hb dbname=snapp port=5432 sslmode=disable"
	Db, Err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if Err != nil {
		log.Fatal(Err)
	} else {
		fmt.Println("Successfully connected to database :)")
	}

	Db.AutoMigrate(&models.TempRule{})
}
