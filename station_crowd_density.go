package ltadatamall

import (
	"errors"
	"time"
)

var AllTrainLines = []string{
	"CCL", "CEL", "CGL", "DTL", "EWL", "NEL",
	"NSL", "BPL", "SLRT", "PLRT", "TEL",
}

type RealTimeStationCrowdDensity struct {
	Station    string    `json:"Station"`
	StartTime  time.Time `json:"StartTime"`
	EndTime    time.Time `json:"EndTime"`
	CrowdLevel string    `json:"CrowdLevel"`
}

type RealTimeStationCrowdDensityResponse struct {
	Metadata                    string `json:"odata.metadata"`
	TrainLine                   string
	RealTimeStationCrowdDensity []RealTimeStationCrowdDensity `json:"value"`
}

type AllRealTimeStationCrowdDensityResponse struct {
	Lines []RealTimeStationCrowdDensityResponse
}

type CrowdLevelInterval struct {
	Start      time.Time `json:"Start"`
	CrowdLevel string    `json:"CrowdLevel"`
}
type ForecastStationCrowdDensity struct {
	Station  string               `json:"Station"`
	Interval []CrowdLevelInterval `json:"Interval"`
}

type DatedForecastStationCrowdDensity struct {
	Date     time.Time                     `json:"Date"`
	Stations []ForecastStationCrowdDensity `json:"Stations"`
}
type ForecastStationCrowdDensityResponse struct {
	Metadata                    string `json:"odata.metadata"`
	TrainLine                   string
	ForecastStationCrowdDensity []DatedForecastStationCrowdDensity `json:"value"`
}
type AllForecastStationCrowdDensityResponse struct {
	Lines []ForecastStationCrowdDensityResponse
}

func GetStationCrowdDensityRealTimeByLine(apiClient *APIClient, line string) (RealTimeStationCrowdDensityResponse, error) {
	var result RealTimeStationCrowdDensityResponse
	endpoint := "PCDRealTime?TrainLine=" + line

	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return RealTimeStationCrowdDensityResponse{}, err
	}

	if len(result.RealTimeStationCrowdDensity) == 0 {
		return RealTimeStationCrowdDensityResponse{}, errors.New("no station crowd densities available")
	}
	result.TrainLine = line
	return result, nil

}

func GetAllStationCrowdDensityRealTime(apiClient *APIClient) (AllRealTimeStationCrowdDensityResponse, error) {
	var result AllRealTimeStationCrowdDensityResponse
	for line := range AllTrainLines {
		res, err := GetStationCrowdDensityRealTimeByLine(apiClient, AllTrainLines[line])
		if err != nil {
			return AllRealTimeStationCrowdDensityResponse{}, err
		}

		result.Lines = append(result.Lines, res)
	}
	return result, nil
}

func GetStationCrowdDensityForecastByLine(apiClient *APIClient, line string) (ForecastStationCrowdDensityResponse, error) {
	var result ForecastStationCrowdDensityResponse
	endpoint := "PCDForecast?TrainLine=" + line

	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return ForecastStationCrowdDensityResponse{}, err
	}

	if len(result.ForecastStationCrowdDensity) == 0 {
		return ForecastStationCrowdDensityResponse{}, errors.New("no station crowd densities available")
	}
	result.TrainLine = line
	return result, nil

}

func GetAllStationCrowdDensityForecast(apiClient *APIClient) (AllForecastStationCrowdDensityResponse, error) {
	var result AllForecastStationCrowdDensityResponse
	for line := range AllTrainLines {
		res, err := GetStationCrowdDensityForecastByLine(apiClient, AllTrainLines[line])
		if err != nil {
			return AllForecastStationCrowdDensityResponse{}, err
		}

		result.Lines = append(result.Lines, res)
	}
	return result, nil
}
