package ltadatamall

type FacilityMaintenance struct {
	Line        string `json:"Line"`
	StationCode string `json:"StationCode"`
	StationName string `json:"StationName"`
	LiftID      string `json:"LiftID"`
	LiftDesc    string `json:"LiftDesc"`
}

type AllFacilityMaintenanceResponse struct {
	Metadata         string                `json:"odata.metadata"`
	LiftMaintenances []FacilityMaintenance `json:"value"`
}

func GetAllFacilityMaintenance(apiClient *APIClient) (AllFacilityMaintenanceResponse, error) {
	var result AllFacilityMaintenanceResponse
	endpoint := "v2/FacilitiesMaintenance"

	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllFacilityMaintenanceResponse{}, err
	}

	if len(result.LiftMaintenances) == 0 {
		return AllFacilityMaintenanceResponse{}, nil
	}

	return result, nil
}
