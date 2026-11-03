package ltadatamall

import "testing"

func TestGetAllRoadOpenings(t *testing.T) {
	result, err := GetAllRoadOpenings(testClient)
	if err != nil {
		t.Errorf("GetAllRoadOpenings returned an error: %v", err)
	}
	if len(result.RoadOpenings) == 0 {
		t.Errorf("GetAllRoadOpenings returned no data, expected at least one planned road opening")
	}
}

func TestGetRoadOpeningsPaginated_NoMoreData(t *testing.T) {
	_, err := GetRoadOpeningsPaginated(testClient, 1000000)
	if err == nil {
		t.Fatalf("Expected error when no more planned road openings are available")
	}
}

func TestGetAllRoadWorks(t *testing.T) {
	result, err := GetAllRoadWorks(testClient)
	if err != nil {
		t.Errorf("GetAllRoadWorks returned an error: %v", err)
	}
	if len(result.RoadWorks) == 0 {
		t.Errorf("GetAllRoadWorks returned no data, expected at least one approved road work")
	}
}

func TestGetRoadWorksPaginated_NoMoreData(t *testing.T) {
	_, err := GetRoadWorksPaginated(testClient, 1000000)
	if err == nil {
		t.Fatalf("Expected error when no more approved road works are available")
	}
}
