package ltadatamall

import "strconv"

type AllBicycleParkingResponse struct {
	OdataMetadata string           `json:"odata.metadata"`
	Value         []BicycleParking `json:"value"`
}

type BicycleParking struct {
	Description      string  `json:"Description"`
	Latitude         float64 `json:"Latitude"`
	Longitude        float64 `json:"Longitude"`
	RackType         string  `json:"RackType"`
	RackCount        int     `json:"RackCount"`
	ShelterIndicator string  `json:"ShelterIndicator"`
}

func GetBicycleParking(apiClient *APIClient, latitude float64, longitude float64, distance float64) (AllBicycleParkingResponse, error) {
	var result AllBicycleParkingResponse
	endpoint := "BicycleParkingv2?Lat=" + strconv.FormatFloat(latitude, 'f', 6, 64) +
		"&Long=" + strconv.FormatFloat(longitude, 'f', 6, 64) +
		"&Dist=" + strconv.FormatFloat(distance, 'f', 4, 64)
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllBicycleParkingResponse{}, err
	}
	return result, nil
}
