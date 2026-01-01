package ltadatamall

import "testing"

func TestGetAllTrafficIncidents(t *testing.T) {
	result, err := GetAllTrafficIncidents(testClient)
	if err != nil {
		t.Errorf("GetAllTrafficIncidents returned an error: %v", err)
	}
	if len(result.TrafficIncidents) == 0 {
		t.Logf("GetAllTrafficIncidents returned no data, this may be expected if there are currently no traffic incidents")
	}
}
