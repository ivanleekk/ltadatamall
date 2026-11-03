package ltadatamall

import (
	"errors"
	"strconv"
)

type AllVMSResponse struct {
	Metadata string       `json:"odata.metadata"`
	VMS      []VMSMessage `json:"value"`
}

type VMSMessage struct {
	EquipmentID string  `json:"EquipmentID"`
	Latitude    float64 `json:"Latitude"`
	Longitude   float64 `json:"Longitude"`
	Message     string  `json:"Message"`
}

func GetVMSPaginated(apiClient *APIClient, skip int) (AllVMSResponse, error) {
	var result AllVMSResponse
	endpoint := "VMS?$skip=" + strconv.Itoa(skip)

	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllVMSResponse{}, err
	}

	if len(result.VMS) == 0 {
		return AllVMSResponse{}, errors.New("no VMS/EMAS messages available")
	}

	return result, nil
}

func GetAllVMS(apiClient *APIClient) (AllVMSResponse, error) {
	var messages []VMSMessage
	// Keep fetching until all records are retrieved
	errorCount := 0
	pagination := 0
	var res AllVMSResponse
	for errorCount < 1 {
		page, err := GetVMSPaginated(apiClient, pagination)
		if err != nil {
			errorCount++
			break
		}
		res = page
		pagination += 500
		messages = append(messages, page.VMS...)
	}
	result := AllVMSResponse{
		VMS:      messages,
		Metadata: res.Metadata,
	}
	return result, nil
}
