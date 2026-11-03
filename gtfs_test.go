package ltadatamall

import "testing"

func TestGetGTFSScheduleTrain(t *testing.T) {
	result, err := GetGTFSScheduleTrain(testClient)
	if err != nil {
		t.Errorf("GetGTFSScheduleTrain returned an error: %v", err)
	}
	if len(result.Value) == 0 {
		t.Errorf("GetGTFSScheduleTrain returned no data, expected a download link")
	}
}

func TestGetGTFSRealTimeTrainServiceAlerts(t *testing.T) {
	result, err := GetGTFSRealTimeTrainServiceAlerts(testClient)
	if err != nil {
		t.Errorf("GetGTFSRealTimeTrainServiceAlerts returned an error: %v", err)
	}
	if len(result.Value) == 0 {
		t.Errorf("GetGTFSRealTimeTrainServiceAlerts returned no data, expected a download link")
	}
}

func TestGetGTFSRealTimeTrainTripUpdates(t *testing.T) {
	result, err := GetGTFSRealTimeTrainTripUpdates(testClient)
	if err != nil {
		t.Errorf("GetGTFSRealTimeTrainTripUpdates returned an error: %v", err)
	}
	if len(result.Value) == 0 {
		t.Errorf("GetGTFSRealTimeTrainTripUpdates returned no data, expected a download link")
	}
}
