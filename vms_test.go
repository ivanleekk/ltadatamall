package ltadatamall

import "testing"

func TestGetAllVMS(t *testing.T) {
	result, err := GetAllVMS(testClient)
	if err != nil {
		t.Errorf("GetAllVMS returned an error: %v", err)
	}
	if len(result.VMS) == 0 {
		t.Logf("GetAllVMS returned no data, this may be expected if no EMAS signboards are currently displaying a message")
	}
}
