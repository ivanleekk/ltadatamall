package ltadatamall

import "testing"

func TestGetAllPlannedBusRoutes(t *testing.T) {
	result, err := GetAllPlannedBusRoutes(testClient)
	if err != nil {
		t.Errorf("GetAllPlannedBusRoutes returned an error: %v", err)
	}
	if len(result.PlannedBusRoutes) == 0 {
		t.Logf("GetAllPlannedBusRoutes returned no data, this may be expected if there are no planned bus route changes")
	}
}
