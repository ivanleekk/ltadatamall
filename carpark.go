package ltadatamall

import (
	"errors"
	"strconv"
)

type AllCarparkAvailabilityResponse struct {
	Metadata string    `json:"odata.metadata"`
	Carparks []Carpark `json:"value"`
}

type Carpark struct {
	CarParkID     string `json:"CarParkID"`
	Area          string `json:"Area"`
	Development   string `json:"Development"`
	Location      string `json:"Location"`
	AvailableLots int    `json:"AvailableLots"`
	LotType       string `json:"LotType"`
	Agency        string `json:"Agency"`
}

type EVID struct {
	Id     string `json:"id"`
	EvCpId string `json:"evCpId"`
	Status string `json:"status"`
}

type PlugType struct {
	PlugType      string `json:"plugType"`
	PowerRating   string `json:"powerRating"`
	ChargingSpeed string `json:"chargingSpeed"`
	Price         string `json:"price"`
	PriceType     string `json:"priceType"`
	EvIds         []EVID `json:"evIds"`
}

type ChargingPoint struct {
	Status         string     `json:"status"`
	OperatingHours string     `json:"operatingHours"`
	Operator       string     `json:"operator"`
	Position       string     `json:"position"`
	Name           string     `json:"name"`
	Id             string     `json:"id"`
	PlugTypes      []PlugType `json:"plugTypes"`
}

type EVChargerLocationData struct {
	Address        string          `json:"address"`
	Name           string          `json:"name"`
	Longitude      float64         `json:"longitude"`
	Latitude       float64         `json:"latitude"`
	LocationId     string          `json:"locationId"`
	Status         string          `json:"status"`
	ChargingPoints []ChargingPoint `json:"chargingPoints"`
}

type EVChargerLocationValue struct {
	EVChargerLocationData []EVChargerLocationData `json:"evLocationsData"`
}

type AllEVChargerLocationResponse struct {
	Value EVChargerLocationValue `json:"value"`
}

func GetAllCarparkAvailability(apiClient *APIClient) (AllCarparkAvailabilityResponse, error) {
	var carparkAvailabilities []Carpark
	// Keep fetching until all records are retrieved
	errorCount := 0
	pagination := 0
	var res AllCarparkAvailabilityResponse
	for errorCount < 1 {
		res, err := GetCarparkAvailabilityPaginated(apiClient, pagination)
		if err != nil {
			errorCount++
			break
		}
		pagination += 500
		carparkAvailabilities = append(carparkAvailabilities, res.Carparks...)
	}
	result := AllCarparkAvailabilityResponse{
		Carparks: carparkAvailabilities,
		Metadata: res.Metadata,
	}
	return result, nil
}

func GetCarparkAvailabilityPaginated(apiClient *APIClient, skip int) (AllCarparkAvailabilityResponse, error) {
	var result AllCarparkAvailabilityResponse
	endpoint := "CarParkAvailabilityv2?$skip=" + strconv.Itoa(skip)

	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllCarparkAvailabilityResponse{}, err
	}

	if len(result.Carparks) == 0 {
		return AllCarparkAvailabilityResponse{}, errors.New("no bus services available")
	}

	return result, nil

}

func GetEVChargersByPostalCode(apiClient *APIClient, postalCode int) (AllEVChargerLocationResponse, error) {
	var result AllEVChargerLocationResponse
	endpoint := "EVChargingPoints?PostalCode=" + strconv.Itoa(postalCode)
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllEVChargerLocationResponse{}, err
	}

	if len(result.Value.EVChargerLocationData) == 0 {
		return AllEVChargerLocationResponse{}, errors.New("no EV charger locations available for the given postal code")
	}

	return result, nil
}
