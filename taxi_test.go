package ltadatamall

import "testing"

func TestGetTaxiAvailability(t *testing.T) {
	response, err := GetTaxiAvailability(testClient)
	if err != nil {
		t.Fatalf("Error calling GetTaxiAvailability: %v", err)
	}

	if len(response.Taxis) <= 500 {
		t.Errorf("Expected more than 500 Taxis in response")
	}
}

func TestGetTaxiAvailabilityPaginated(t *testing.T) {
	response, err := GetTaxiAvailabilityPaginated(testClient, 500)
	if err != nil {
		t.Fatalf("Error calling GetBusStopsPaginated: %v", err)
	}
	if len(response.Taxis) == 0 {
		t.Errorf("Expected non-empty Taxis in response")
	}
}

func TestGetTaxiAvailabilityPaginated_NoMoreData(t *testing.T) {
	_, err := GetTaxiAvailabilityPaginated(testClient, 1000000)
	if err == nil {
		t.Fatalf("Expected error when no more taxis are available")
	}
}
