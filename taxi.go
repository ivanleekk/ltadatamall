package ltadatamall

import "fmt"

type TaxiCoordinates struct {
	Latitude  float32 `json:"Latitude"`
	Longitude float32 `json:"Longitude"`
}

type TaxiAvailabilityResponse struct {
	Metadata string            `json:"odata.metadata"`
	Taxis    []TaxiCoordinates `json:"value"`
}

type TaxiStand struct {
	TaxiCode  string  `json:"TaxiCode"`
	Latitude  float32 `json:"Latitude"`
	Longitude float32 `json:"Longitude"`
	Bfa       string  `json:"Bfa"`
	Ownership string  `json:"Ownership"`
	Type      string  `json:"Type"`
	Name      string  `json:"Name"`
}

type TaxiStandResponse struct {
	Metadata   string      `json:"odata.metadata"`
	TaxiStands []TaxiStand `json:"value"`
}

func GetAllTaxiAvailability(apiClient *APIClient) (TaxiAvailabilityResponse, error) {
	var taxis []TaxiCoordinates
	// Keep fetching until all records are retrieved
	errorCount := 0
	pagination := 0
	var res TaxiAvailabilityResponse
	for errorCount < 1 {
		res, err := GetTaxiAvailabilityPaginated(apiClient, pagination)
		if err != nil {
			errorCount++
			break
		}
		pagination += 500
		taxis = append(taxis, res.Taxis...)
	}
	result := TaxiAvailabilityResponse{
		Taxis:    taxis,
		Metadata: res.Metadata,
	}
	return result, nil
}

func GetTaxiAvailabilityPaginated(apiClient *APIClient, skip int) (TaxiAvailabilityResponse, error) {
	var result TaxiAvailabilityResponse
	endpoint := "Taxi-Availability?$skip=" + fmt.Sprint(skip)
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return TaxiAvailabilityResponse{}, err
	}

	if len(result.Taxis) == 0 {
		return TaxiAvailabilityResponse{}, fmt.Errorf("no more taxi availability data")
	}

	return result, nil
}

func GetAllTaxiStands(apiClient *APIClient) (TaxiStandResponse, error) {
	var taxiStands []TaxiStand
	// Keep fetching until all records are retrieved
	errorCount := 0
	pagination := 0
	var res TaxiStandResponse
	for errorCount < 1 {
		res, err := GetTaxiStandsPaginated(apiClient, pagination)
		if err != nil {
			errorCount++
			break
		}
		pagination += 500
		taxiStands = append(taxiStands, res.TaxiStands...)
	}
	result := TaxiStandResponse{
		TaxiStands: taxiStands,
		Metadata:   res.Metadata,
	}
	return result, nil
}

func GetTaxiStandsPaginated(apiClient *APIClient, skip int) (TaxiStandResponse, error) {
	var result TaxiStandResponse
	endpoint := "TaxiStands?$skip=" + fmt.Sprint(skip)
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return TaxiStandResponse{}, err
	}

	if len(result.TaxiStands) == 0 {
		return TaxiStandResponse{}, fmt.Errorf("no more taxi stand data")
	}

	return result, nil
}
