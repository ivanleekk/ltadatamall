package ltadatamall

func GetAllTrafficImages(apiClient *APIClient) (AllTrafficImagesResponse, error) {
	var result AllTrafficImagesResponse
	endpoint := "Traffic-Imagesv2"
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllTrafficImagesResponse{}, err
	}

	if len(result.Value) == 0 {
		return AllTrafficImagesResponse{}, nil
	}

	return result, nil
}

type AllTrafficImagesResponse struct {
	Metadata string              `json:"odata.metadata"`
	Value    []TrafficImageEntry `json:"value"`
}

type TrafficImageEntry struct {
	CameraID  string  `json:"CameraID"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	ImageLink string  `json:"ImageLink"`
}
