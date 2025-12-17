package ltadatamall

import "testing"

func TestGetTaxiAvailability(t *testing.T) {
	response, err := GetTaxiAvailability(testClient)
	if err != nil {
		t.Fatalf("Error calling GetTaxiAvailability: %v", err)
	}

	if len(response.Taxis) == 0 {
		t.Errorf("Expected non-empty Taxis in response")
	}
}
