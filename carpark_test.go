package ltadatamall

import (
	"testing"
)

func TestGetAllCarparkAvailability(t *testing.T) {
	response, err := GetAllCarparkAvailability(testClient)
	if err != nil {
		t.Fatalf("Error calling GetAllCarparkAvailability: %v", err)
	}

	if len(response.Carparks) <= 500 {
		t.Errorf("Expected more than 500 CarparkAvailability in response")
	}
}

func TestGetCarparkAvailabilityPaginated(t *testing.T) {
	response, err := GetCarparkAvailabilityPaginated(testClient, 500)
	if err != nil {
		t.Fatalf("Error calling GetCarparkAvailabilityPaginated: %v", err)
	}
	if len(response.Carparks) == 0 {
		t.Errorf("Expected non-empty CarparkAvailability in response")
	}
}

func TestGetCarparkAvailabilityPaginated_NoMoreData(t *testing.T) {
	_, err := GetCarparkAvailabilityPaginated(testClient, 1000000)
	if err == nil {
		t.Fatalf("Expected error when no more carparks are available")
	}
}
