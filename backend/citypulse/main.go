package main

import (
	"fmt"
	"os"

	"backend/gtfs" // adjust to your module name

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("could not load .env:", err)
	}

	lat := os.Getenv("CITY_LAT")
	lon := os.Getenv("CITY_LON")

	data, err := gtfs.Fetch(lat, lon)
	if err != nil {
		fmt.Println("fetch error:", err)
		return
	}

	fmt.Println(data)
}
