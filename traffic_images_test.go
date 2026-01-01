package ltadatamall

import (
	"testing"
)

func TestGetAllTrafficImages(t *testing.T) {
	result, err := GetAllTrafficImages(testClient)
	if err != nil {
		t.Errorf("GetAllTrafficImages returned an error: %v", err)
	}
	if len(result.Value) == 0 {
		t.Errorf("GetAllTrafficImages returned no data, expected at least one traffic image")
	}
}
