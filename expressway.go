package ltadatamall

import "errors"

type AllExpresswayEstimatedTravelTimes struct {
	Metadata             string                           `json:"odata.metadata"`
	EstimatedTravelTimes []ExpresswayEstimatedTravelTimes `json:"value"`
}

type ExpresswayEstimatedTravelTimes struct {
	Name        string `json:"Name"`
	Direction   int    `json:"Direction"`
	FarEndPoint string `json:"FarEndPoint"`
	StartPoint  string `json:"StartPoint"`
	EndPoint    string `json:"EndPoint"`
	EstTime     int    `json:"EstTime"`
}

func GetAllExpresswayEstimatedTravelTimes(apiClient *APIClient) (AllExpresswayEstimatedTravelTimes, error) {
	var result AllExpresswayEstimatedTravelTimes
	endpoint := "EstTravelTimes"

	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllExpresswayEstimatedTravelTimes{}, err
	}

	if len(result.EstimatedTravelTimes) == 0 {
		return AllExpresswayEstimatedTravelTimes{}, errors.New("no estimated travel times found")
	}

	return result, nil
}
