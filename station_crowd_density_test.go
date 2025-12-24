package ltadatamall

import (
	"log"
	"testing"
)

func TestGetAllStationCrowdDensityRealTime(t *testing.T) {
	response, err := GetAllStationCrowdDensityRealTime(testClient)
	if err != nil {
		t.Fatalf("Error calling GetAllStationCrowdDensityRealTime: %v", err)
	}
	log.Default().Println("AllStationCrowdDensityRealTime", response)
	if len(response.Lines) != 11 {
		t.Errorf("Expected Exactly 11 Lines of RealTime Station Crowd Density in response")
	}
}

func TestGetStationCrowdDensityRealTimeByLine(t *testing.T) {
	response, err := GetStationCrowdDensityRealTimeByLine(testClient, "TEL")
	if err != nil {
		t.Fatalf("Error calling GetStationCrowdDensityRealTimeByLine: %v", err)
	}

	if len(response.RealTimeStationCrowdDensity) < 27 {
		t.Errorf("Expected at least 27 RealTime Station Crowd Density in response")
	}
}

func TestGetAllStationCrowdDensityForecast(t *testing.T) {
	response, err := GetAllStationCrowdDensityForecast(testClient)
	if err != nil {
		t.Fatalf("Error calling GetAllStationCrowdDensityForecast: %v", err)
	}
	log.Default().Println("GetAllStationCrowdDensityForecast", response)
	if len(response.Lines) != 11 {
		t.Errorf("Expected Exactly 11 Lines of Forecast Station Crowd Density in response")
	}
}

func TestGetStationCrowdDensityForecastByLine(t *testing.T) {
	response, err := GetStationCrowdDensityForecastByLine(testClient, "TEL")
	if err != nil {
		t.Fatalf("Error calling GetStationCrowdDensityForecastByLine: %v", err)
	}

	if len(response.ForecastStationCrowdDensity[0].Stations) < 27 {
		t.Errorf("Expected at least 27 Forecast Station Crowd Density in response")
	}
}
