package ltadatamall

import (
	"errors"
	"strconv"
)

type AllFaultyTrafficLightsResponse struct {
	Metadata            string               `json:"odata.metadata"`
	FaultyTrafficLights []FaultyTrafficLight `json:"value"`
}

type FaultyTrafficLight struct {
	AlarmID     string `json:"AlarmID"`
	NodeID      string `json:"NodeID"`
	Type        int    `json:"Type"`
	StartDate   string `json:"StartDate"`
	EndDate     string `json:"EndDate"`
	XCoordinate string `json:"XCoordinate"`
	YCoordinate string `json:"YCoordinate"`
	Message     string `json:"Message"`
}

func GetFaultyTrafficLightsPaginated(apiClient *APIClient, skip int) (AllFaultyTrafficLightsResponse, error) {
	var result AllFaultyTrafficLightsResponse
	endpoint := "v2/FaultyTrafficLights?$skip=" + strconv.Itoa(skip)

	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllFaultyTrafficLightsResponse{}, err
	}

	if len(result.FaultyTrafficLights) == 0 {
		return AllFaultyTrafficLightsResponse{}, errors.New("no faulty traffic lights available")
	}

	return result, nil
}

func GetAllFaultyTrafficLights(apiClient *APIClient) (AllFaultyTrafficLightsResponse, error) {
	var faultyTrafficLights []FaultyTrafficLight
	// Keep fetching until all records are retrieved
	errorCount := 0
	pagination := 0
	var res AllFaultyTrafficLightsResponse
	for errorCount < 1 {
		page, err := GetFaultyTrafficLightsPaginated(apiClient, pagination)
		if err != nil {
			errorCount++
			break
		}
		res = page
		pagination += 500
		faultyTrafficLights = append(faultyTrafficLights, page.FaultyTrafficLights...)
	}
	result := AllFaultyTrafficLightsResponse{
		FaultyTrafficLights: faultyTrafficLights,
		Metadata:            res.Metadata,
	}
	return result, nil
}
