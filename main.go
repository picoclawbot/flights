package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	departure := flag.String("departure", "", "Departure airport code (e.g., JFK)")
	destination := flag.String("destination", "", "Destination airport code (e.g., CDG)")
	startDate := flag.String("start_date", "", "Departure date (YYYY-MM-DD)")
	endDate := flag.String("end_date", "", "Return date (YYYY-MM-DD)")
	cabinClass := flag.String("cabin", "Economy", "Cabin class (Economy, Business, etc.)")

	flag.Parse()

	if *departure == "" || *destination == "" || *startDate == "" || *endDate == "" {
		fmt.Println("Usage: flights_go -departure JFK -destination CDG -start_date 2024-05-01 -end_date 2024-05-10 [-cabin Business]")
		os.Exit(1)
	}

	fmt.Printf("Searching for %s flights from %s to %s from %s to %s...\n", *cabinClass, *departure, *destination, *startDate, *endDate)

	flights, err := SearchFlights(*departure, *destination, *startDate, *endDate, *cabinClass)
	if err != nil {
		log.Fatalf("Error searching flights: %v", err)
	}

	if len(flights) == 0 {
		fmt.Println("No flights found.")
		return
	}

	fmt.Printf("Found %d flight options:\n", len(flights))
	
	var cheapest *FlightInfo
	minPrice := 999999

	for i, f := range flights {
		fmt.Printf("[%d] %s: %s (%s, %s) - %s to %s\n", i+1, f.Airline, f.Price, f.Duration, f.Stops, f.DepartureTime, f.ArrivalTime)
		
		priceVal := parsePrice(f.Price)
		if priceVal > 0 && priceVal < minPrice {
			minPrice = priceVal
			cheapest = &flights[i]
		}
	}

	if cheapest != nil {
		fmt.Printf("\n--- Cheapest Option ---\n")
		fmt.Printf("Airline: %s\n", cheapest.Airline)
		fmt.Printf("Price: %s\n", cheapest.Price)
		fmt.Printf("Duration: %s\n", cheapest.Duration)
		fmt.Printf("Stops: %s\n", cheapest.Stops)
		fmt.Printf("Times: %s - %s\n", cheapest.DepartureTime, cheapest.ArrivalTime)
		fmt.Printf("URL: %s\n", cheapest.URL)
	}
}
