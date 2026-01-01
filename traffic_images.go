package ltadatamall

type AllTrafficImagesResponse struct {
	Metadata      string              `json:"odata.metadata"`
	TrafficImages []TrafficImageEntry `json:"trafficimages"`
}

type TrafficImageEntry struct {
	CameraID  string  `json:"CameraID"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	ImageLink string  `json:"ImageLink"`
}

func GetAllTrafficImages(apiClient *APIClient) (AllTrafficImagesResponse, error) {
	var result AllTrafficImagesResponse
	endpoint := "Traffic-Imagesv2"
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllTrafficImagesResponse{}, err
	}

	if len(result.TrafficImages) == 0 {
		return AllTrafficImagesResponse{}, nil
	}

	return result, nil
}
