package ltadatamall

type TrainServiceAlertAffectedSegment struct {
	Line                string `json:"Line"`
	Direction           string `json:"Direction"`
	Stations            string `json:"Stations"`
	FreePublicBus       string `json:"FreePublicBus"`
	FreeMRTShuttle      string `json:"FreeMRTShuttle"`
	MRTShuttleDirection string `json:"MRTShuttleDirection"`
}

type TrainServiceAlertMessage struct {
	Content     string `json:"Content"`
	CreatedDate string `json:"CreatedDate"`
}

type TrainServiceAlert struct {
	Status           int                                `json:"Status"`
	AffectedSegments []TrainServiceAlertAffectedSegment `json:"AffectedSegments"`
	Message          []TrainServiceAlertMessage         `json:"Message"`
}

type TrainServiceAlertsResponse struct {
	Metadata string            `json:"odata.metadata"`
	Value    TrainServiceAlert `json:"value"`
}

func GetTrainServiceAlerts(apiClient *APIClient) (TrainServiceAlertsResponse, error) {
	var result TrainServiceAlertsResponse
	endpoint := "TrainServiceAlerts"
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return TrainServiceAlertsResponse{}, err
	}

	return result, nil
}
