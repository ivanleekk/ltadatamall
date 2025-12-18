package ltadatamall

import (
	"errors"
	"strconv"
)

type BusService struct {
	ServiceNo       string `json:"ServiceNo"`
	Operator        string `json:"Operator"`
	Direction       int    `json:"Direction"`
	Category        string `json:"Category"`
	OriginCode      string `json:"OriginCode"`
	DestinationCode string `json:"DestinationCode"`
	AM_Peak_Freq    string `json:"AM_Peak_Freq"`
	AM_Offpeak_Freq string `json:"AM_Offpeak_Freq"`
	PM_Peak_Freq    string `json:"PM_Peak_Freq"`
	PM_Offpeak_Freq string `json:"PM_Offpeak_Freq"`
	LoopDesc        string `json:"LoopDesc"`
}

type AllBusServiceResponse struct {
	Metadata    string       `json:"odata.metadata"`
	BusServices []BusService `json:"value"`
}

type BusRoute struct {
	ServiceNo    string  `json:"ServiceNo"`
	Operator     string  `json:"Operator"`
	Direction    int     `json:"Direction"`
	StopSequence int     `json:"StopSequence"`
	BusStopCode  string  `json:"BusStopCode"`
	Distance     float32 `json:"Distance"`
	WD_FirstBus  string  `json:"WD_FirstBus"`
	WD_LastBus   string  `json:"WD_LastBus"`
	SAT_FirstBus string  `json:"SAT_FirstBus"`
	SAT_LastBus  string  `json:"SAT_LastBus"`
	SUN_FirstBus string  `json:"SUN_FirstBus"`
	SUN_LastBus  string  `json:"SUN_LastBus"`
}

type AllBusRouteResponse struct {
	Metadata  string     `json:"odata.metadata"`
	BusRoutes []BusRoute `json:"value"`
}

type BusStop struct {
	BusStopCode string  `json:"BusStopCode"`
	RoadName    string  `json:"RoadName"`
	Description string  `json:"Description"`
	Latitude    float32 `json:"Latitude"`
	Longitude   float32 `json:"Longitude"`
}

type AllBusStopResponse struct {
	Metadata string    `json:"odata.metadata"`
	BusStops []BusStop `json:"value"`
}

type NextBus struct {
	OriginCode       string `json:"OriginCode"`
	DestinationCode  string `json:"DestinationCode"`
	EstimatedArrival string `json:"EstimatedArrival"`
	Monitored        int    `json:"Monitored"`
	Latitude         string `json:"Latitude"`
	Longitude        string `json:"Longitude"`
	VisitNumber      string `json:"VisitNumber"`
	Load             string `json:"Load"`
	Feature          string `json:"Feature"`
	Type             string `json:"Type"`
}

type BusArrivalService struct {
	ServiceNo string  `json:"ServiceNo"`
	Operator  string  `json:"Operator"`
	NextBus   NextBus `json:"NextBus"`
	NextBus2  NextBus `json:"NextBus2"`
	NextBus3  NextBus `json:"NextBus3"`
}

type BusArrivalResponse struct {
	Metadata    string              `json:"odata.metadata"`
	BusStopCode string              `json:"BusStopCode"`
	Services    []BusArrivalService `json:"Services"`
}

func GetAllBusServices(apiClient *APIClient) (AllBusServiceResponse, error) {
	var busServices []BusService
	// Keep fetching until all records are retrieved
	errorCount := 0
	pagination := 0
	var res AllBusServiceResponse
	for errorCount < 1 {
		res, err := GetBusServicesPaginated(apiClient, len(busServices))
		if err != nil {
			errorCount++
			break
		}
		pagination += 500
		busServices = append(busServices, res.BusServices...)
	}
	result := AllBusServiceResponse{
		BusServices: busServices,
		Metadata:    res.Metadata,
	}
	return result, nil
}

func GetBusServicesPaginated(apiClient *APIClient, skip int) (AllBusServiceResponse, error) {
	var result AllBusServiceResponse
	endpoint := "BusServices?$skip=" + strconv.Itoa(skip)

	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllBusServiceResponse{}, err
	}

	if len(result.BusServices) == 0 {
		return AllBusServiceResponse{}, errors.New("no bus services available")
	}

	return result, nil

}
func GetAllBusRoutes(apiClient *APIClient) (AllBusRouteResponse, error) {
	var result AllBusRouteResponse
	if err := apiClient.getJSON("BusRoutes", &result); err != nil {
		return AllBusRouteResponse{}, err
	}

	return result, nil
}

func GetAllBusStops(apiClient *APIClient) (AllBusStopResponse, error) {
	var result AllBusStopResponse
	if err := apiClient.getJSON("BusStops", &result); err != nil {
		return AllBusStopResponse{}, err
	}

	return result, nil
}

func GetBusArrivalAtBusStop(apiClient *APIClient, busStopCode string) (BusArrivalResponse, error) {
	var result BusArrivalResponse
	endpoint := "v3/BusArrival?BusStopCode=" + busStopCode
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return BusArrivalResponse{}, err
	}

	return result, nil
}

func GetBusArrivalAtBusStopAndService(apiClient *APIClient, busStopCode, service string) (BusArrivalResponse, error) {
	var result BusArrivalResponse
	endpoint := "v3/BusArrival?BusStopCode=" + busStopCode + "&ServiceNo=" + service
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return BusArrivalResponse{}, err
	}

	return result, nil
}
