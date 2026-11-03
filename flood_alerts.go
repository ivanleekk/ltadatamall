package ltadatamall

type AllFloodAlertsResponse struct {
	Metadata    string       `json:"odata.metadata"`
	FloodAlerts []FloodAlert `json:"value"`
}

type FloodAlert struct {
	AlertID      string `json:"alertId"`
	DateTime     string `json:"dateTime"`
	MsgType      string `json:"msgType"`
	Event        string `json:"event"`
	ResponseType string `json:"responseType"`
	Urgency      string `json:"urgency"`
	Severity     string `json:"severity"`
	Expires      string `json:"expires"`
	SenderName   string `json:"senderName"`
	Headline     string `json:"headline"`
	Description  string `json:"description"`
	Instruction  string `json:"instruction"`
	AreaDesc     string `json:"areaDesc"`
	Circle       string `json:"circle"`
	Status       string `json:"status"`
}

func GetAllFloodAlerts(apiClient *APIClient) (AllFloodAlertsResponse, error) {
	var result AllFloodAlertsResponse
	endpoint := "PubFloodAlerts"
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllFloodAlertsResponse{}, err
	}

	if len(result.FloodAlerts) == 0 {
		return AllFloodAlertsResponse{}, nil
	}

	return result, nil
}
