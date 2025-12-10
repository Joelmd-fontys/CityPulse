package fetcher

// Used AI to better udnerstand the JSON
type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Coordinates struct {
	Coordinate []Coordinate `json:"coordinate"`
}

type FlowSegmentData struct {
	CurrentSpeed       int         `json:"currentSpeed"`
	FreeFlowSpeed      int         `json:"freeFlowSpeed"`
	CurrentTravelTime  int         `json:"currentTravelTime"`
	FreeFlowTravelTime int         `json:"freeFlowTravelTime"`
	Confidence         float64     `json:"confidence"`
	Coordinates        Coordinates `json:"coordinates"`
}

type TrafficAPIResponse struct {
	FlowSegmentData FlowSegmentData `json:"flowSegmentData"`
}

func (f FlowSegmentData) Congestion() float64 {
	if f.FreeFlowSpeed == 0 {
		return 0
	}
	raw := 1.0 - float64(f.CurrentSpeed)/float64(f.FreeFlowSpeed)
	if raw < 0 {
		raw = 0
	}
	if raw > 1 {
		raw = 1
	}
	return raw
}
