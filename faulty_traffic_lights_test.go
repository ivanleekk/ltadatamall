package ltadatamall

import "testing"

func TestGetAllFaultyTrafficLights(t *testing.T) {
	result, err := GetAllFaultyTrafficLights(testClient)
	if err != nil {
		t.Errorf("GetAllFaultyTrafficLights returned an error: %v", err)
	}
	if len(result.FaultyTrafficLights) == 0 {
		t.Logf("GetAllFaultyTrafficLights returned no data, this may be expected if there are currently no faulty traffic lights")
	}
}
