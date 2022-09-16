package validations

import (
	"encoding/csv"
	"example/snapp/models"
	"fmt"
	"os"
	"strings"
)

var cities []string
var Cities = map[string]bool{}
var airlines []string
var Airlines = map[string]bool{}
var agencies []string
var Agencies = map[string]bool{}
var suppliers []string
var Suppliers = map[string]bool{}

func PreValidation() {
	// city
	var cityRecords [][]string

	cityFile, err := os.Open("city.csv")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer cityFile.Close()

	cityReader := csv.NewReader(cityFile)
	cityRecords, _ = cityReader.ReadAll()

	for i, j := range cityRecords {
		cities = append(cities, j[2])
		fmt.Println(cities[i])
		Cities[cities[i]] = true
	}
	// airline
	var airlineRecords [][]string

	airlineFile, err := os.Open("airline.csv")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer airlineFile.Close()

	airlineReader := csv.NewReader(airlineFile)
	airlineRecords, _ = airlineReader.ReadAll()

	for i, j := range airlineRecords {
		airlines = append(airlines, strings.ToLower(j[0]))

		fmt.Println(airlines[i])
		Airlines[airlines[i]] = true
	}
	// agency
	var agencyRecords [][]string

	agencyFile, err := os.Open("agency.csv")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer agencyFile.Close()

	agencyReader := csv.NewReader(agencyFile)
	agencyRecords, _ = agencyReader.ReadAll()

	for i, j := range agencyRecords {
		agencies = append(agencies, j[2])
		fmt.Println(agencies[i])
		Agencies[agencies[i]] = true
	}
	// supplier
	var supplierRecords [][]string

	supplierFile, err := os.Open("supplier.csv")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer agencyFile.Close()

	supplierReader := csv.NewReader(supplierFile)
	supplierRecords, _ = supplierReader.ReadAll()

	for i, j := range supplierRecords {
		suppliers = append(suppliers, j[2])
		fmt.Println(suppliers[i])
		Suppliers[suppliers[i]] = true
	}
}

func ValidateRule(t models.Rule) (isValid bool) {
	return true
}

func ValidateTicket(t models.Ticket) (isValid bool) {
	return true
}
