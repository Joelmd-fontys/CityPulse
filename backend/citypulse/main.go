package main

import (
	"backend/fetcher"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("could not load .env:", err)
	}

	lat := os.Getenv("CITY_LAT")
	lon := os.Getenv("CITY_LON")

	data, err := fetcher.Fetch(lat, lon)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Speed:", data.FlowSegmentData.CurrentSpeed)
	fmt.Println("Free flow:", data.FlowSegmentData.FreeFlowSpeed)
	fmt.Println("First coordinate:", data.FlowSegmentData.Coordinates.Coordinate[0])
}
