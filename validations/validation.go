package validations

import (
	"context"
	"encoding/csv"
	"example/snapp/databases"
	"example/snapp/models"
	"fmt"
	"os"
	"strings"
)

var ctx context.Context

var Cities = map[string]bool{}
var Airlines = map[string]bool{}
var Agencies = map[string]bool{}
var Suppliers = map[string]bool{}

func PreValidation() {
	ctx = context.Background()
	PreValidationCity()
	PreValidationAirline()
	PreValidationAgency()
	PreValidationSupplier()
}

func PreValidationCity() {
	var cities []string
	var err error
	if databases.GetValue("ValidCityListCreated") == "true" {
		cities, err = databases.Client.LRange(ctx, "Valid:City", 0, -1).Result()
		if err != nil {
			fmt.Printf("PreValidationCity: %s\n", err.Error())
		}
	} else if databases.GetValue("ValidCityTableCreated") == "true" {

	} else {
		var cityRecords [][]string
		cityFile, err := os.Open("city.csv")
		if err != nil {
			fmt.Printf("PreValidationCity: %s\n", err.Error())
			return
		}
		cityReader := csv.NewReader(cityFile)
		cityRecords, _ = cityReader.ReadAll()
		cityFile.Close()

		for _, j := range cityRecords {
			cities = append(cities, j[2])
			databases.CreateValidCityTable(j[2])
			databases.SetValue("ValidCityTableCreated", "true")
			databases.CreateValidCityList(j[2])
			databases.SetValue("ValidCityListCreated", "true")
		}
	}

	for _, j := range cities {
		Cities[j] = true
	}
}

func PreValidationAirline() {
	var airlines []string
	var err error

	if databases.GetValue("ValidAirlineListCreated") == "true" {
		airlines, err = databases.Client.LRange(ctx, "Valid:Airline", 0, -1).Result()
		if err != nil {
			fmt.Printf("PreValidationAirline: %s\n", err.Error())
		}

	} else if databases.GetValue("ValidAirlineTableCreated") == "true" {

	} else {
		var airlineRecords [][]string
		airlineFile, err := os.Open("airline.csv")
		if err != nil {
			fmt.Printf("PreValidationAirline: %s\n", err.Error())
			return
		}
		airlineReader := csv.NewReader(airlineFile)
		airlineRecords, _ = airlineReader.ReadAll()
		airlineFile.Close()

		for _, j := range airlineRecords {
			airlines = append(airlines, strings.ToLower(j[0]))
			databases.CreateValidAirlineTable(strings.ToLower(j[0]))
			databases.SetValue("ValidAirlineTableCreated", "true")
			databases.CreateValidAirlineList(strings.ToLower(j[0]))
			databases.SetValue("ValidAirlineListCreated", "true")
		}
	}

	for _, j := range airlines {
		Airlines[j] = true
	}

}

func PreValidationAgency() {
	var agencies []string
	var err error

	if databases.GetValue("ValidAgencyListCreated") == "true" {
		agencies, err = databases.Client.LRange(ctx, "Valid:Agency", 0, -1).Result()
		if err != nil {
			fmt.Printf("PreValidationAgency: %s\n", err.Error())
		}

	} else if databases.GetValue("ValidAgencyTableCreated") == "true" {

	} else {
		var agencyRecords [][]string
		agencyFile, err := os.Open("agency.csv")
		if err != nil {
			fmt.Printf("PreValidationAgency: %s\n", err.Error())
			return
		}
		agencyReader := csv.NewReader(agencyFile)
		agencyRecords, _ = agencyReader.ReadAll()
		agencyFile.Close()

		for _, j := range agencyRecords {
			agencies = append(agencies, j[2])
			databases.CreateValidAgencyTable(j[2])
			databases.SetValue("ValidAgencyTableCreated", "true")
			databases.CreateValidAgencyList(j[2])
			databases.SetValue("ValidAgencyListCreated", "true")
		}
	}

	for _, j := range agencies {
		Agencies[j] = true
	}
}

func PreValidationSupplier() {
	var suppliers []string
	var err error

	if databases.GetValue("ValidSupplierListCreated") == "true" {
		suppliers, err = databases.Client.LRange(ctx, "Valid:Supplier", 0, -1).Result()
		if err != nil {
			fmt.Printf("PreValidationSupplier: %s\n", err.Error())
		}

	} else if databases.GetValue("ValidSupplierTableCreated") == "true" {

	} else {
		var supplierRecords [][]string
		supplierFile, err := os.Open("supplier.csv")
		if err != nil {
			fmt.Printf("PreValidationSupplier: %s\n", err.Error())
			return
		}
		supplierReader := csv.NewReader(supplierFile)
		supplierRecords, _ = supplierReader.ReadAll()
		supplierFile.Close()

		for _, j := range supplierRecords {
			suppliers = append(suppliers, j[2])
			databases.CreateValidSupplierTable(j[2])
			databases.SetValue("ValidSupplierTableCreated", "true")
			databases.CreateValidSupplierList(j[2])
			databases.SetValue("ValidSupplierListCreated", "true")
		}
	}

	for _, j := range suppliers {
		Suppliers[j] = true
	}
}

func ValidateRule(t []models.Rule) bool {
	isValid := true

	for _, i := range t {
		for _, j := range i.Routes {
			if (!Cities[j.Origin] && j.Origin != "") || (!Cities[j.Destination] && j.Destination != "") {
				isValid = false
				break
			}
		}
		if !isValid {
			break
		}
		for _, j := range i.Airlines {
			if !Airlines[j] && j != "" {
				isValid = false
				break
			}
		}
		if !isValid {
			break
		}
		for _, j := range i.Agencies {
			if !Agencies[j] && j != "" {
				isValid = false
				break
			}
		}
		if !isValid {
			break
		}
		for _, j := range i.Suppliers {
			if !Suppliers[j] && j != "" {
				isValid = false
				break
			}
		}
		if !isValid {
			break
		}
	}

	return isValid
}

// func ValidateTicket(t *[]models.Ticket) bool {
// 	return true
// }
