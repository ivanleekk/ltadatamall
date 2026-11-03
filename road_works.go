package ltadatamall

import (
	"errors"
	"strconv"
)

type AllRoadOpeningsResponse struct {
	Metadata     string        `json:"odata.metadata"`
	RoadOpenings []RoadWorkObj `json:"value"`
}

type AllRoadWorksResponse struct {
	Metadata  string        `json:"odata.metadata"`
	RoadWorks []RoadWorkObj `json:"value"`
}

type RoadWorkObj struct {
	EventID   string `json:"EventID"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	SvcDept   string `json:"SvcDept"`
	RoadName  string `json:"RoadName"`
	Other     string `json:"Other"`
}

func GetRoadOpeningsPaginated(apiClient *APIClient, skip int) (AllRoadOpeningsResponse, error) {
	var result AllRoadOpeningsResponse
	endpoint := "RoadOpenings?$skip=" + strconv.Itoa(skip)

	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllRoadOpeningsResponse{}, err
	}

	if len(result.RoadOpenings) == 0 {
		return AllRoadOpeningsResponse{}, errors.New("no planned road openings available")
	}

	return result, nil
}

func GetAllRoadOpenings(apiClient *APIClient) (AllRoadOpeningsResponse, error) {
	var roadOpenings []RoadWorkObj
	// Keep fetching until all records are retrieved
	errorCount := 0
	pagination := 0
	var res AllRoadOpeningsResponse
	for errorCount < 1 {
		page, err := GetRoadOpeningsPaginated(apiClient, pagination)
		if err != nil {
			errorCount++
			break
		}
		res = page
		pagination += 500
		roadOpenings = append(roadOpenings, page.RoadOpenings...)
	}
	result := AllRoadOpeningsResponse{
		RoadOpenings: roadOpenings,
		Metadata:     res.Metadata,
	}
	return result, nil
}

func GetRoadWorksPaginated(apiClient *APIClient, skip int) (AllRoadWorksResponse, error) {
	var result AllRoadWorksResponse
	endpoint := "RoadWorks?$skip=" + strconv.Itoa(skip)

	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return AllRoadWorksResponse{}, err
	}

	if len(result.RoadWorks) == 0 {
		return AllRoadWorksResponse{}, errors.New("no approved road works available")
	}

	return result, nil
}

func GetAllRoadWorks(apiClient *APIClient) (AllRoadWorksResponse, error) {
	var roadWorks []RoadWorkObj
	// Keep fetching until all records are retrieved
	errorCount := 0
	pagination := 0
	var res AllRoadWorksResponse
	for errorCount < 1 {
		page, err := GetRoadWorksPaginated(apiClient, pagination)
		if err != nil {
			errorCount++
			break
		}
		res = page
		pagination += 500
		roadWorks = append(roadWorks, page.RoadWorks...)
	}
	result := AllRoadWorksResponse{
		RoadWorks: roadWorks,
		Metadata:  res.Metadata,
	}
	return result, nil
}
