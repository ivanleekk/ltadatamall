package ltadatamall

import (
	"errors"
	"strconv"
)

type AllPlannedBusRouteResponse struct {
	Metadata         string            `json:"odata.metadata"`
	PlannedBusRoutes []PlannedBusRoute `json:"value"`
}

type PlannedBusRoute struct {
	ServiceNo     string  `json:"ServiceNo"`
	Operator      string  `json:"Operator"`
	Direction     int     `json:"Direction"`
	StopSequence  int     `json:"StopSequence"`
	BusStopCode   string  `json:"BusStopCode"`
	Distance      float32 `json:"Distance"`
	WD_FirstBus   string  `json:"WD_FirstBus"`
	WD_LastBus    string  `json:"WD_LastBus"`
	SAT_FirstBus  string  `json:"SAT_FirstBus"`
	SAT_LastBus   string  `json:"SAT_LastBus"`
	SUN_FirstBus  string  `json:"SUN_FirstBus"`
	SUN_LastBus   string  `json:"SUN_LastBus"`
	EffectiveDate string  `json:"EffectiveDate"`
}

func GetPlannedBusRoutesPaginated(apiClient *APIClient, skip int) (AllPlannedBusRouteResponse, error) {
	var result AllPlannedBusRouteResponse
	endpoint := "PlannedBusRoutes?$skip=" + strconv.Itoa(skip)

	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllPlannedBusRouteResponse{}, err
	}

	if len(result.PlannedBusRoutes) == 0 {
		return AllPlannedBusRouteResponse{}, errors.New("no planned bus routes available")
	}

	return result, nil
}

func GetAllPlannedBusRoutes(apiClient *APIClient) (AllPlannedBusRouteResponse, error) {
	var plannedBusRoutes []PlannedBusRoute
	// Keep fetching until all records are retrieved
	errorCount := 0
	pagination := 0
	var res AllPlannedBusRouteResponse
	for errorCount < 1 {
		page, err := GetPlannedBusRoutesPaginated(apiClient, pagination)
		if err != nil {
			errorCount++
			break
		}
		res = page
		pagination += 500
		plannedBusRoutes = append(plannedBusRoutes, page.PlannedBusRoutes...)
	}
	result := AllPlannedBusRouteResponse{
		PlannedBusRoutes: plannedBusRoutes,
		Metadata:         res.Metadata,
	}
	return result, nil
}
