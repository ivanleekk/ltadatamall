package ltadatamall

import "testing"

func TestGetTrainServiceAlerts(t *testing.T) {
	result, err := GetTrainServiceAlerts(testClient)
	if err != nil {
		t.Errorf("GetTrainServiceAlerts returned an error: %v", err)
	}
	if result.Value.Status == 0 {
		t.Logf("GetTrainServiceAlerts returned no status, this may be expected outside of train operating hours")
	}
}
