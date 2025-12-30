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
