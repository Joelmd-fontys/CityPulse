package main

import (
	"backend/fetcher"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	latStr := os.Getenv("CITY_LAT")
	lonStr := os.Getenv("CITY_LON")
	if latStr == "" || lonStr == "" {
		fmt.Println("CITY_LAT or CITY_LON is missing in .env")
		return
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		fmt.Println("invalid CITY_LAT:", err)
		return
	}
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		fmt.Println("invalid CITY_LON:", err)
		return
	}
	
	grid, err := fetcher.FetchArea(lat, lon)
	if err != nil {
		fmt.Println("error fetching area:", err)
		return
	}

	if err := os.MkdirAll("frontend", 0o755); err != nil {
		fmt.Println("error creating frontend dir:", err)
		return
	}

	outPath := filepath.Join("frontend", "grid.json")
	f, err := os.Create(outPath)
	if err != nil {
		fmt.Println("error creating grid.json:", err)
		return
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(grid); err != nil {
		fmt.Println("error writing JSON:", err)
		return
	}

	fmt.Println("wrote grid to", outPath)
	fmt.Println("grid size:", len(grid), "rows,", len(grid[0]), "cols")
}
