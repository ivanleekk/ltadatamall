package ltadatamall

import (
	"errors"
	"strconv"
)

type AllTrafficSpeedBandsResponse struct {
	Metadata          string             `json:"odata.metadata"`
	LastUpdatedTime   string             `json:"lastUpdatedTime"`
	TrafficSpeedBands []TrafficSpeedBand `json:"value"`
}

type TrafficSpeedBand struct {
	LinkID       string `json:"LinkID"`
	RoadName     string `json:"RoadName"`
	RoadCategory string `json:"RoadCategory"`
	SpeedBand    int    `json:"SpeedBand"`
	MinimumSpeed string `json:"MinimumSpeed"`
	MaximumSpeed string `json:"MaximumSpeed"`
	StartLon     string `json:"StartLon"`
	StartLat     string `json:"StartLat"`
	EndLon       string `json:"EndLon"`
	EndLat       string `json:"EndLat"`
}

func GetTrafficSpeedBandsPaginated(apiClient *APIClient, skip int) (AllTrafficSpeedBandsResponse, error) {
	var result AllTrafficSpeedBandsResponse
	endpoint := "v4/TrafficSpeedBands?$skip=" + strconv.Itoa(skip)
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllTrafficSpeedBandsResponse{}, err
	}

	if len(result.TrafficSpeedBands) == 0 {
		return AllTrafficSpeedBandsResponse{}, errors.New("no TrafficSpeedBands found")
	}
	return result, nil
}

func GetAllTrafficSpeedBands(apiClient *APIClient) (AllTrafficSpeedBandsResponse, error) {
	var trafficSpeedBands []TrafficSpeedBand
	// Keep fetching until all records are retrieved
	errorCount := 0
	pagination := 0
	var res AllTrafficSpeedBandsResponse
	for errorCount < 1 {
		res, err := GetTrafficSpeedBandsPaginated(apiClient, pagination)
		if err != nil {
			errorCount++
			break
		}

		pagination += 500
		trafficSpeedBands = append(trafficSpeedBands, res.TrafficSpeedBands...)
	}
	result := AllTrafficSpeedBandsResponse{
		TrafficSpeedBands: trafficSpeedBands,
		Metadata:          res.Metadata,
		LastUpdatedTime:   res.LastUpdatedTime,
	}
	return result, nil
}
