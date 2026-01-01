package ltadatamall

import (
	"testing"
)

func TestGetTrafficSpeedBandsPaginated(t *testing.T) {
	response, err := GetTrafficSpeedBandsPaginated(testClient, 500)
	if err != nil {
		t.Fatalf("Error calling GetBusServicesPaginated: %v", err)
	}
	if len(response.TrafficSpeedBands) == 0 {
		t.Errorf("Expected non-empty TrafficSpeedBands in response")
	}
}

func TestGetTrafficSpeedBandsPaginated_NoMoreData(t *testing.T) {
	_, err := GetTrafficSpeedBandsPaginated(testClient, 1000000)
	if err == nil {
		t.Fatalf("Expected error when no more traffic speed bands are available")
	}
}

func TestGetAllTrafficSpeedBands(t *testing.T) {
	response, err := GetAllTrafficSpeedBands(testClient)
	if err != nil {
		t.Fatalf("Error calling GetAllTrafficSpeedBands: %v", err)
	}

	if len(response.TrafficSpeedBands) <= 500 {
		t.Errorf("Expected more than 500 TrafficSpeedBands in response")
	}

	t.Logf("Length of TrafficSpeedBands: %d", len(response.TrafficSpeedBands))
}
