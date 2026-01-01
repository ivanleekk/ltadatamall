package ltadatamall

import (
	"testing"
)

func TestFormatYearMonth(t *testing.T) {
	tests := []struct {
		year     int
		month    int
		expected string
	}{
		{2023, 1, "202301"},
		{2023, 10, "202310"},
		{1999, 12, "199912"},
		{2000, 5, "200005"},
	}

	for _, test := range tests {
		result := formatYearMonth(test.year, test.month)
		if result != test.expected {
			t.Errorf("formatYearMonth(%d, %d) = %s; expected %s", test.year, test.month, result, test.expected)
		}
	}
}

func TestGetBusStopsPassengerVolumes(t *testing.T) {
	result, err := GetBusStopsPassengerVolumes(testClient, currentYear, int(currentMonth))
	if err != nil {
		t.Errorf("GetBusStopsPassengerVolumes returned an error: %v", err)
	}
	if len(result.Value) == 0 {
		t.Log("GetBusStopsPassengerVolumes returned no data, which may be expected for the current or future months.")
	}
}

func TestGetOriginDestinationBusStopsPassengerVolumes(t *testing.T) {
	result, err := GetOriginDestinationBusStopsPassengerVolumes(testClient, currentYear, int(currentMonth))
	if err != nil {
		t.Errorf("GetOriginDestinationBusStopsPassengerVolumes returned an error: %v", err)
	}
	if len(result.Value) == 0 {
		t.Log("GetOriginDestinationBusStopsPassengerVolumes returned no data, which may be expected for the current or future months.")
	}
}

func TestGetTrainStationsPassengerVolumes(t *testing.T) {
	result, err := GetTrainStationsPassengerVolumes(testClient, currentYear, int(currentMonth))
	if err != nil {
		t.Errorf("GetTrainStationsPassengerVolumes returned an error: %v", err)
	}
	if len(result.Value) == 0 {
		t.Log("GetTrainStationsPassengerVolumes returned no data, which may be expected for the current or future months.")
	}
}

func TestGetOriginDestinationTrainStationsPassengerVolumes(t *testing.T) {
	result, err := GetOriginDestinationTrainStationsPassengerVolumes(testClient, currentYear, int(currentMonth))
	if err != nil {
		t.Errorf("GetOriginDestinationTrainStationsPassengerVolumes returned an error: %v", err)
	}
	if len(result.Value) == 0 {
		t.Log("GetOriginDestinationTrainStationsPassengerVolumes returned no data, which may be expected for the current or future months.")
	}
}
