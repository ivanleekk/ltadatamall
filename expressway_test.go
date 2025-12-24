package ltadatamall

import (
	"testing"
)

func TestGetAllExpresswayEstimatedTravelTimes(t *testing.T) {
	response, err := GetAllExpresswayEstimatedTravelTimes(testClient)
	if err != nil {
		t.Fatalf("Error calling ExpresswayEstimatedTravelTimes: %v", err)
	}

	if len(response.EstimatedTravelTimes) < 1 {
		t.Errorf("Expected at least 1 ExpresswayEstimatedTravelTimes in response")
	}
}
