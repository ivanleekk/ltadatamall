package ltadatamall

import (
	"testing"
)

func TestGetTrafficFlow(t *testing.T) {
	result, err := GetTrafficFlow(testClient)
	if err != nil {
		t.Errorf("GetTrafficFlow returned an error: %v", err)
	}
	if len(result.Value) == 0 {
		t.Errorf("GetTrafficFlow returned empty value")
	}
}
