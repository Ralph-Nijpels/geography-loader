package main

// Data-loader puts the Country, Region and Airport information from CSV files
// in the MongoDB.
//
// Files are retrieved from ourairports.com/data/xx.csv
//
// Note: it is written quite sloppily:
// - error logging is not implemented

import (
	"fmt"
	"log"

	application "github.com/ralph-nijpels/geography-application/v2"
	countries "github.com/ralph-nijpels/geography-countries/v2"
	airports "github.com/ralph-nijpels/geography-airports/v2"
)

func main() {

	fmt.Println("Initializing..")
	context, err := application.CreateAppContext()
	if err != nil {
		log.Fatal(err)
	}
	defer context.Destroy()

	fmt.Println("Loading countries..")
	countries, err := countries.NewCountries(context)
	if err != nil {
		log.Fatal(err)
	}
	err = countries.RetrieveFromURL()
	if err != nil {
		log.Fatal(err)
	}
	err = countries.ImportCSV()
	if err != nil {
		log.Fatal(err)
	}

	// check
	_, err = countries.GetByCountryCode("US")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Loading regions..")
	regions := countries.NewRegions()
	err = regions.RetrieveFromURL()
	if err != nil {
		log.Fatal(err)
	}
	err = regions.ImportCSV()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Loading airports..")
	airports,err := airports.NewAirports(context, countries)
	err = airports.RetrieveFromURL()
	if err != nil {
		log.Fatal(err)
	}
	err = airports.ImportCSV()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Loading runways..")
	runways := airports.NewRunways()
	err = runways.RetrieveFromURL()
	if err != nil {
		log.Fatal(err)
	}
	err = runways.ImportCSV()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Loading frequencies..")
	frequencies := airports.NewFrequencies()
	err = frequencies.RetrieveFromURL()
	if err != nil {
		log.Fatal(err)
	}
	err = frequencies.ImportCSV()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data loaded.")
}
