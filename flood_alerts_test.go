package ltadatamall

import "testing"

func TestGetAllFloodAlerts(t *testing.T) {
	result, err := GetAllFloodAlerts(testClient)
	if err != nil {
		t.Errorf("GetAllFloodAlerts returned an error: %v", err)
	}
	if len(result.FloodAlerts) == 0 {
		t.Logf("GetAllFloodAlerts returned no data, this may be expected if there are currently no flood alerts")
	}
}
