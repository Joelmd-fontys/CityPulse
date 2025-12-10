package gtfs

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Fetch(lat, lon string) (string, error) {
	apiKey := os.Getenv("TOMTOM_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("TOMTOM_API_KEY is empty")
	}

	url := fmt.Sprintf(
		"https://api.tomtom.com/traffic/services/4/flowSegmentData/absolute/10/json?point=%s,%s&unit=KMPH&key=%s",
		lat, lon, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("non-200 status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
