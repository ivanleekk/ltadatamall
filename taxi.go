package ltadatamall

type TaxiCoordinates struct {
	Latitude  float32 `json:"Latitude"`
	Longitude float32 `json:"Longitude"`
}

type TaxiAvailabilityResponse struct {
	Metadata string            `json:"odata.metadata"`
	Taxis    []TaxiCoordinates `json:"value"`
}

func GetTaxiAvailability(apiClient *APIClient) (TaxiAvailabilityResponse, error) {
	var result TaxiAvailabilityResponse
	if err := apiClient.getJSON("Taxi-Availability", &result); err != nil {
		return TaxiAvailabilityResponse{}, err
	}

	return result, nil
}
