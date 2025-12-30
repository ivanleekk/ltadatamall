package ltadatamall

import (
	"testing"
)

// TestGetGeospatialLayer_AllEndpoints programmatically tests all geospatial layer enums
// by iterating over the geospatialEndpoint map. This ensures any new enums added
// to the map are automatically tested.
func TestGetGeospatialLayer_AllEndpoints(t *testing.T) {
	// Iterate over all geospatial layer types defined in the endpoint map
	for layerType, endpointName := range geospatialEndpoint {
		// Capture range variables for parallel subtests
		layerType := layerType
		endpointName := endpointName

		t.Run(endpointName, func(t *testing.T) {
			result, err := GetGeospatialLayer(testClient, layerType)
			if err != nil {
				t.Errorf("GeospatialLayer %s returned an error: %v", endpointName, err)
			}
			if len(result.Value) == 0 {
				t.Errorf("GeospatialLayer %s returned empty value", endpointName)
			}
		})
	}
}

// TestGetGeospatialLayer_AllEnumsHaveEndpoints verifies that all defined enum constants
// have corresponding entries in the geospatialEndpoint map
func TestGetGeospatialLayer_AllEnumsHaveEndpoints(t *testing.T) {
	// List of all geospatial layer enum values
	allEnums := []geospatialLayerType{
		GeospatialLayerArrowMarking,
		GeospatialLayerBollard,
		GeospatialLayerBusStopLocation,
		GeospatialLayerControlBox,
		GeospatialLayerConvexMirror,
		GeospatialLayerCoveredLinkWay,
		GeospatialLayerCyclingPath,
		GeospatialLayerDetectorLoop,
		GeospatialLayerERPGantry,
		GeospatialLayerFootpath,
		GeospatialLayerGuardRail,
		GeospatialLayerKerbLine,
		GeospatialLayerLampPost,
		GeospatialLayerLaneMarking,
		GeospatialLayerParkingStandardsZone,
		GeospatialLayerPassengerPickupBay,
		GeospatialLayerPedestrainOverheadbridge_UnderPass,
		GeospatialLayerRailConstruction,
		GeospatialLayerRailing,
		GeospatialLayerRetainingWall,
		GeospatialLayerRoadCrossing,
		GeospatialLayerRoadHump,
		GeospatialLayerRoadSectionLine,
		GeospatialLayerSchoolZone,
		GeospatialLayerSilverZone,
		GeospatialLayerSpeedRegulatingStrip,
		GeospatialLayerStreetPaint,
		GeospatialLayerTaxiStand,
		GeospatialLayerTrafficLight,
		GeospatialLayerTrafficSign,
		GeospatialLayerTrainStation,
		GeospatialLayerTrainStationExit,
		GeospatialLayerVehicularBridge_Flyover_Underpass,
		GeospatialLayerWordMarking,
	}

	for _, enumVal := range allEnums {
		if _, exists := geospatialEndpoint[enumVal]; !exists {
			t.Errorf("Enum value %d has no corresponding endpoint in geospatialEndpoint map", enumVal)
		}
	}

	// Also verify count matches
	if len(allEnums) != len(geospatialEndpoint) {
		t.Errorf("Mismatch: %d enum values but %d endpoints in map", len(allEnums), len(geospatialEndpoint))
	}
}

// TestGetAllGeospatialLayers tests that GetAllGeospatialLayers returns all layers
// with proper success/failure tracking
func TestGetAllGeospatialLayers(t *testing.T) {
	response := GetAllGeospatialLayers(testClient)

	// Verify total count matches the number of defined endpoints
	expectedTotal := len(geospatialEndpoint)
	if response.TotalCount != expectedTotal {
		t.Errorf("Expected TotalCount to be %d, got %d", expectedTotal, response.TotalCount)
	}

	// Verify the number of layers returned matches total
	if len(response.Layers) != expectedTotal {
		t.Errorf("Expected %d layers, got %d", expectedTotal, len(response.Layers))
	}

	// Verify success + failure = total
	if response.SuccessCount+response.FailureCount != response.TotalCount {
		t.Errorf("SuccessCount (%d) + FailureCount (%d) should equal TotalCount (%d)",
			response.SuccessCount, response.FailureCount, response.TotalCount)
	}

	// Log results for visibility
	t.Logf("GetAllGeospatialLayers: %d success, %d failures, %d total",
		response.SuccessCount, response.FailureCount, response.TotalCount)

	// Check each layer result
	for _, layer := range response.Layers {
		if layer.Name == "" {
			t.Errorf("Layer has empty name")
		}
		if layer.Error != nil {
			// Log failures but don't fail the test - some endpoints may be temporarily unavailable
			t.Logf("Warning: Layer %s failed with error: %v", layer.Name, layer.Error)
		} else {
			if len(layer.Response.Value) == 0 {
				t.Errorf("Layer %s has no value but no error", layer.Name)
			}
		}
	}

	// Expect most endpoints to succeed (allow some failures for API issues)
	minSuccessRate := 0.9 // At least 90% should succeed
	actualSuccessRate := float64(response.SuccessCount) / float64(response.TotalCount)
	if actualSuccessRate < minSuccessRate {
		t.Errorf("Success rate too low: %.1f%% (expected at least %.1f%%)",
			actualSuccessRate*100, minSuccessRate*100)
	}
}
