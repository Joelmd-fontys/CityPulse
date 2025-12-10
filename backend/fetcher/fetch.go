package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
)

func Fetch(lat, lon string) (*TrafficAPIResponse, error) {
	apiKey := os.Getenv("TOMTOM_API_KEY")

	url := fmt.Sprintf(
		"https://api.tomtom.com/traffic/services/4/flowSegmentData/absolute/10/json?point=%s,%s&unit=KMPH&key=%s",
		lat, lon, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("non-200: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data TrafficAPIResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func FetchArea(centerLat, centerLon float64) ([][]float64, error) {
	const halfSizeMeters = 500.0 // 1 km total → 500 m each direction
	const stepMeters = 200.0     // ~200 m step

	latStep := stepMeters / 111000.0
	latRadius := halfSizeMeters / 111000.0

	lonMeterPerDeg := 111000.0 * math.Cos(centerLat*math.Pi/180.0)
	lonStep := stepMeters / lonMeterPerDeg
	lonRadius := halfSizeMeters / lonMeterPerDeg

	var grid [][]float64

	for dLat := -latRadius; dLat <= latRadius+1e-9; dLat += latStep {
		var row []float64
		for dLon := -lonRadius; dLon <= lonRadius+1e-9; dLon += lonStep {
			lat := centerLat + dLat
			lon := centerLon + dLon

			resp, err := Fetch(
				fmt.Sprintf("%.6f", lat),
				fmt.Sprintf("%.6f", lon),
			)
			if err != nil {
				return nil, err
			}

			row = append(row, resp.FlowSegmentData.Congestion())
		}
		grid = append(grid, row)
	}

	return grid, nil
}
