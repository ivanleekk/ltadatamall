package ltadatamall

import (
	"strconv"
)

type AllPassengerVolumeResponse struct {
	Metadata string                `json:"odata.metadata"`
	Value    []PassengerVolumeLink `json:"value"`
	Type     string
	Year     int
	Month    int
}

type PassengerVolumeLink struct {
	Link string `json:"Link"`
}

func formatYearMonth(year int, month int) string {
	var monthStr string
	if month < 10 {
		monthStr = "0" + strconv.Itoa(month)
	} else {
		monthStr = strconv.Itoa(month)
	}
	return strconv.Itoa(year) + monthStr
}

func GetBusStopsPassengerVolumes(apiClient *APIClient, year int, month int) (AllPassengerVolumeResponse, error) {
	var result AllPassengerVolumeResponse
	endpoint := "PV/Bus?Date=" + formatYearMonth(year, month)

	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllPassengerVolumeResponse{}, err
	}

	if len(result.Value) == 0 {
		return AllPassengerVolumeResponse{}, nil
	}
	result.Type = "Bus"
	result.Year = year
	result.Month = month
	return result, nil
}

func GetOriginDestinationBusStopsPassengerVolumes(apiClient *APIClient, year int, month int) (AllPassengerVolumeResponse, error) {
	var result AllPassengerVolumeResponse
	endpoint := "PV/ODBus?Date=" + formatYearMonth(year, month)
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllPassengerVolumeResponse{}, err
	}

	if len(result.Value) == 0 {
		return AllPassengerVolumeResponse{}, nil
	}
	result.Type = "ODBus"
	result.Year = year
	result.Month = month
	return result, nil
}

func GetTrainStationsPassengerVolumes(apiClient *APIClient, year int, month int) (AllPassengerVolumeResponse, error) {
	var result AllPassengerVolumeResponse
	endpoint := "PV/Train?Date=" + formatYearMonth(year, month)
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllPassengerVolumeResponse{}, err
	}

	if len(result.Value) == 0 {
		return AllPassengerVolumeResponse{}, nil
	}
	result.Type = "Train"
	result.Year = year
	result.Month = month
	return result, nil
}

func GetOriginDestinationTrainStationsPassengerVolumes(apiClient *APIClient, year int, month int) (AllPassengerVolumeResponse, error) {
	var result AllPassengerVolumeResponse
	endpoint := "PV/ODTrain?Date=" + formatYearMonth(year, month)
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllPassengerVolumeResponse{}, err
	}

	if len(result.Value) == 0 {
		return AllPassengerVolumeResponse{}, nil
	}
	result.Type = "ODTrain"
	result.Year = year
	result.Month = month
	return result, nil
}
