package ltadatamall

type AllTrafficIncidentResponse struct {
	OdataMetadata    string            `json:"odata.metadata"`
	TrafficIncidents []TrafficIncident `json:"value"`
}

type TrafficIncident struct {
	Type      string  `json:"Type"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Message   string  `json:"Message"`
}

func GetAllTrafficIncidents(apiClient *APIClient) (AllTrafficIncidentResponse, error) {
	var result AllTrafficIncidentResponse
	endpoint := "TrafficIncidents"
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllTrafficIncidentResponse{}, err
	}

	if len(result.TrafficIncidents) == 0 {
		return AllTrafficIncidentResponse{}, nil
	}

	return result, nil
}
