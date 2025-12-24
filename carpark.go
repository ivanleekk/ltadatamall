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
