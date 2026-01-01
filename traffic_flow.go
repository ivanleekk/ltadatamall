package ltadatamall

type AllTrafficFlowResponse struct {
	Metadata    string           `json:"odata.metadata"`
	TrafficFlow []TrafficFlowObj `json:"value"`
}

type TrafficFlowObj struct {
	Link string `json:"Link"`
}

func GetTrafficFlow(apiClient *APIClient) (AllTrafficFlowResponse, error) {
	var result AllTrafficFlowResponse
	endpoint := "TrafficFlow"
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllTrafficFlowResponse{}, err
	}

	if len(result.TrafficFlow) == 0 {
		return AllTrafficFlowResponse{}, nil
	}

	return result, nil
}
