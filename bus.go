package ltadatamall

import (
	"encoding/json"
	"errors"
	"net/http"
)

type BusService struct {
	ServiceNo       string `json:"ServiceNo"`
	Operator        string `json:"Operator"`
	Direction       int    `json:"Direction"`
	Category        string `json:"Category"`
	OriginCode      string `json:"OriginCode"`
	DestinationCode string `json:"DestinationCode"`
	AM_Peak_Freq    string `json:"AM_Peak_Freq"`
	AM_Offpeak_Freq string `json:"AM_Offpeak_Freq"`
	PM_Peak_Freq    string `json:"PM_Peak_Freq"`
	PM_Offpeak_Freq string `json:"PM_Offpeak_Freq"`
	LoopDesc        string `json:"LoopDesc"`
}

type AllBusServiceResponse struct {
	Metadata    string       `json:"odata.metadata"`
	BusServices []BusService `json:"value"`
}

func GetAllBusService(client *APIClient) (AllBusServiceResponse, error) {
	req, err := http.NewRequest("GET", client.baseURL+"BusServices", nil)
	if err != nil {
		return AllBusServiceResponse{}, errors.New("Error creating the request")
	}
	req.Header.Add("AccountKey", client.accountKey)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return AllBusServiceResponse{}, err
	}
	defer resp.Body.Close()

	var result AllBusServiceResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return AllBusServiceResponse{}, err
	}

	return result, nil
}
