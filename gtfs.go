package ltadatamall

type DownloadLink struct {
	Link string `json:"Link"`
}

type AllGTFSScheduleTrainResponse struct {
	Metadata string         `json:"odata.metadata"`
	Value    []DownloadLink `json:"value"`
}

type AllGTFSRealTimeTrainServiceAlertsResponse struct {
	Metadata string         `json:"odata.metadata"`
	Value    []DownloadLink `json:"value"`
}

type AllGTFSRealTimeTrainTripUpdatesResponse struct {
	Metadata string         `json:"odata.metadata"`
	Value    []DownloadLink `json:"value"`
}

// GetGTFSScheduleTrain returns the download link for the static GTFS Schedule
// (Train) ZIP file. The link expires after 15 minutes.
func GetGTFSScheduleTrain(apiClient *APIClient) (AllGTFSScheduleTrainResponse, error) {
	var result AllGTFSScheduleTrainResponse
	endpoint := "GTFSScheduleTrain"
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllGTFSScheduleTrainResponse{}, err
	}

	if len(result.Value) == 0 {
		return AllGTFSScheduleTrainResponse{}, nil
	}

	return result, nil
}

// GetGTFSRealTimeTrainServiceAlerts returns the download link for the GTFS
// Realtime (Train Service Alerts) protobuf feed. The link expires after 15 minutes.
func GetGTFSRealTimeTrainServiceAlerts(apiClient *APIClient) (AllGTFSRealTimeTrainServiceAlertsResponse, error) {
	var result AllGTFSRealTimeTrainServiceAlertsResponse
	endpoint := "GTFSRealTimeTrainServiceAlerts"
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllGTFSRealTimeTrainServiceAlertsResponse{}, err
	}

	if len(result.Value) == 0 {
		return AllGTFSRealTimeTrainServiceAlertsResponse{}, nil
	}

	return result, nil
}

// GetGTFSRealTimeTrainTripUpdates returns the download link for the GTFS
// Realtime (Train Trip Updates - Disruption) protobuf feed. The link expires
// after 15 minutes.
func GetGTFSRealTimeTrainTripUpdates(apiClient *APIClient) (AllGTFSRealTimeTrainTripUpdatesResponse, error) {
	var result AllGTFSRealTimeTrainTripUpdatesResponse
	endpoint := "GTFSRealtimeTrainTripUpdates"
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllGTFSRealTimeTrainTripUpdatesResponse{}, err
	}

	if len(result.Value) == 0 {
		return AllGTFSRealTimeTrainTripUpdatesResponse{}, nil
	}

	return result, nil
}
