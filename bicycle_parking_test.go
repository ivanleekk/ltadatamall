package ltadatamall

import "testing"

func TestGetBicycleParking(t *testing.T) {
	result, err := GetBicycleParking(testClient, 1.364897, 103.766094, 0.456)
	if err != nil {
		t.Errorf("GetBicycleParking returned an error: %v", err)
	}
	if len(result.Value) == 0 {
		t.Errorf("GetBicycleParking returned no data, expected at least one bicycle parking location")
	}

}
