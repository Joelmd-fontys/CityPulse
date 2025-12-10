package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
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
