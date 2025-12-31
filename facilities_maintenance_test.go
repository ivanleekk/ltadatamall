package ltadatamall

import (
	"testing"
)

func TestGetAllFacilityMaintenance(t *testing.T) {
	_, err := GetAllFacilityMaintenance(testClient)
	if err != nil {
		t.Errorf("GetAllFacilityMaintenance returned an error: %v", err)
	}
}
