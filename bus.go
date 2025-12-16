package ltadatamall

import (
	"encoding/json"
	"errors"
	"net/http"
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

func GetAllBusServices(client *APIClient) (AllBusServiceResponse, error) {
	req, err := http.NewRequest("GET", client.baseURL+"BusServices", nil)
	if err != nil {
		return AllBusServiceResponse{}, errors.New("error creating the request")
	}
	req.Header.Add("AccountKey", client.accountKey)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return AllBusServiceResponse{}, err
	}
	defer resp.Body.Close()

	var result AllBusServiceResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return AllBusServiceResponse{}, err
	}

	return result, nil
}

func GetAllBusRoutes(client *APIClient) (AllBusRouteResponse, error) {
	req, err := http.NewRequest("GET", client.baseURL+"BusRoutes", nil)
	if err != nil {
		return AllBusRouteResponse{}, errors.New("error creating the request")
	}
	req.Header.Add("AccountKey", client.accountKey)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return AllBusRouteResponse{}, err
	}
	defer resp.Body.Close()

	var result AllBusRouteResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return AllBusRouteResponse{}, err
	}

	return result, nil
}

func GetAllBusStops(client *APIClient) (AllBusStopResponse, error) {
	req, err := http.NewRequest("GET", client.baseURL+"BusStops", nil)
	if err != nil {
		return AllBusStopResponse{}, errors.New("error creating the request")
	}
	req.Header.Add("AccountKey", client.accountKey)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return AllBusStopResponse{}, err
	}
	defer resp.Body.Close()

	var result AllBusStopResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return AllBusStopResponse{}, err
	}

	return result, nil
}
