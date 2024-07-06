package main

import (
	"fmt"
	"os"
	"time"

	"./warehouse"
)

func main() {
	s, err := warehouse.NewServerFromFile("warehouse_pings.csv")
	if err != nil {
		fmt.Printf("Failed to create warehouse server: %v\n", err)
		os.Exit(1)
	}
	fmt.Print("~~~WarehouseServer is initialized.\n\n")

	fmt.Printf("Average Speeds: %v\n\n", s.AverageSpeeds())

	startTime := time.Unix(1553273158, 0)
	fmt.Printf("The 3 most traveled vehicles since %d are:\n%v\n", startTime.Unix(), s.MostTraveledSince(3, startTime))

	fmt.Printf("Vehicles possibly damaged:\n%v\n", s.CheckForDamage())
}
