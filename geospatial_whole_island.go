package ltadatamall

import (
	"errors"
)

type geospatialLayerType int

const (
	GeospatialLayerArrowMarking = iota
	GeospatialLayerBollard
	GeospatialLayerBusStopLocation
	GeospatialLayerControlBox
	GeospatialLayerConvexMirror
	GeospatialLayerCoveredLinkWay
	GeospatialLayerCyclingPath
	GeospatialLayerDetectorLoop
	GeospatialLayerERPGantry
	GeospatialLayerFootpath
	GeospatialLayerGuardRail
	GeospatialLayerKerbLine
	GeospatialLayerLampPost
	GeospatialLayerLaneMarking
	GeospatialLayerParkingStandardsZone
	GeospatialLayerPassengerPickupBay
	GeospatialLayerPedestrainOverheadbridge_UnderPass
	GeospatialLayerRailConstruction
	GeospatialLayerRailing
	GeospatialLayerRetainingWall
	GeospatialLayerRoadCrossing
	GeospatialLayerRoadHump
	GeospatialLayerRoadSectionLine
	GeospatialLayerSchoolZone
	GeospatialLayerSilverZone
	GeospatialLayerSpeedRegulatingStrip
	GeospatialLayerStreetPaint
	GeospatialLayerTaxiStand
	GeospatialLayerTrafficLight
	GeospatialLayerTrafficSign
	GeospatialLayerTrainStation
	GeospatialLayerTrainStationExit
	GeospatialLayerVehicularBridge_Flyover_Underpass
	GeospatialLayerWordMarking
)

var geospatialEndpoint = map[geospatialLayerType]string{
	GeospatialLayerArrowMarking:                       "ArrowMarking",
	GeospatialLayerBollard:                            "Bollard",
	GeospatialLayerBusStopLocation:                    "BusStopLocation",
	GeospatialLayerControlBox:                         "ControlBox",
	GeospatialLayerConvexMirror:                       "ConvexMirror",
	GeospatialLayerCoveredLinkWay:                     "CoveredLinkWay",
	GeospatialLayerCyclingPath:                        "CyclingPath",
	GeospatialLayerDetectorLoop:                       "DetectorLoop",
	GeospatialLayerERPGantry:                          "ERPGantry",
	GeospatialLayerFootpath:                           "Footpath",
	GeospatialLayerGuardRail:                          "GuardRail",
	GeospatialLayerKerbLine:                           "KerbLine",
	GeospatialLayerLampPost:                           "LampPost",
	GeospatialLayerLaneMarking:                        "LaneMarking",
	GeospatialLayerParkingStandardsZone:               "ParkingStandardsZone",
	GeospatialLayerPassengerPickupBay:                 "PassengerPickupBay",
	GeospatialLayerPedestrainOverheadbridge_UnderPass: "PedestrainOverheadbridge_UnderPass",
	GeospatialLayerRailConstruction:                   "RailConstruction",
	GeospatialLayerRailing:                            "Railing",
	GeospatialLayerRetainingWall:                      "RetainingWall",
	GeospatialLayerRoadCrossing:                       "RoadCrossing",
	GeospatialLayerRoadHump:                           "RoadHump",
	GeospatialLayerRoadSectionLine:                    "RoadSectionLine",
	GeospatialLayerSchoolZone:                         "SchoolZone",
	GeospatialLayerSilverZone:                         "SilverZone",
	GeospatialLayerSpeedRegulatingStrip:               "SpeedRegulatingStrip",
	GeospatialLayerStreetPaint:                        "StreetPaint",
	GeospatialLayerTaxiStand:                          "TaxiStand",
	GeospatialLayerTrafficLight:                       "TrafficLight",
	GeospatialLayerTrafficSign:                        "TrafficSign",
	GeospatialLayerTrainStation:                       "TrainStation",
	GeospatialLayerTrainStationExit:                   "TrainStationExit",
	GeospatialLayerVehicularBridge_Flyover_Underpass:  "VehicularBridge_Flyover_Underpass",
	GeospatialLayerWordMarking:                        "WordMarking",
}

type GeospatialLink struct {
	Link string `json:"Link"`
}
type RawGeospatialLayerResponse struct {
	Metadata string           `json:"odata.metadata"`
	Value    []GeospatialLink `json:"value"`
}

// Function to call the raw api and just return the output without modification
func GetGeospatialLayer(apiClient *APIClient, layerType geospatialLayerType) (RawGeospatialLayerResponse, error) {
	var result RawGeospatialLayerResponse
	endpoint := "GeospatialWholeIsland?ID=" + geospatialEndpoint[layerType]
	if err := apiClient.getJSON(endpoint, &result); err != nil {
		return RawGeospatialLayerResponse{}, err
	}

	if len(result.Value) == 0 {
		return RawGeospatialLayerResponse{}, errors.New("no geospatial layer at available")
	}

	return result, nil
}
