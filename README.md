# LTA DataMall Go Client

This repository provides a Go client for accessing the LTA DataMall APIs, which offer various transportation-related data services in Singapore.

> [!NOTE]
> All API calls require an API key from LTA DataMall. Please refer to the [LTA DataMall website](https://datamall.lta.gov.sg/content/datamall/en/request-for-api.html) for more information on obtaining an API key.
> API Reference: [LTA DataMall API Documentation](https://datamall.lta.gov.sg/content/dam/datamall/datasets/LTA_DataMall_API_User_Guide.pdf)

A sample `.env` file is provided in `.env.example`. Rename it to `.env` and replace the placeholder values with your actual API key.

## Supported APIs

Every API documented in the LTA DataMall API User Guide (version 6.11, 30 Oct 2026) is covered.

| Guide section | API | Functions |
| --- | --- | --- |
| 2.1 | Bus Arrival | `GetBusArrivalAtBusStop`, `GetBusArrivalAtBusStopAndService` |
| 2.2 | Bus Services | `GetAllBusServices`, `GetBusServicesPaginated`, `GetBusServiceByServiceNo` |
| 2.3 | Bus Routes | `GetAllBusRoutes`, `GetBusRoutesPaginated` |
| 2.4 | Bus Stops | `GetAllBusStops`, `GetBusStopsPaginated`, `GetBusStopByBusStopCode` |
| 2.5 | Passenger Volume by Bus Stops | `GetBusStopsPassengerVolumes` |
| 2.6 | Passenger Volume by Origin Destination Bus Stops | `GetOriginDestinationBusStopsPassengerVolumes` |
| 2.7 | Passenger Volume by Origin Destination Train Stations | `GetOriginDestinationTrainStationsPassengerVolumes` |
| 2.8 | Passenger Volume by Train Stations | `GetTrainStationsPassengerVolumes` |
| 2.9 | Taxi Availability | `GetAllTaxiAvailability`, `GetTaxiAvailabilityPaginated` |
| 2.10 | Taxi Stands | `GetAllTaxiStands` |
| 2.11 | Train Service Alerts | `GetTrainServiceAlerts` |
| 2.12 | Carpark Availability | `GetAllCarparkAvailability`, `GetCarparkAvailabilityPaginated` |
| 2.13 | Estimated Travel Times | `GetAllExpresswayEstimatedTravelTimes` |
| 2.14 | Faulty Traffic Lights | `GetAllFaultyTrafficLights`, `GetFaultyTrafficLightsPaginated` |
| 2.15 | Planned Road Openings | `GetAllRoadOpenings`, `GetRoadOpeningsPaginated` |
| 2.16 | Approved Road Works | `GetAllRoadWorks`, `GetRoadWorksPaginated` |
| 2.17 | Traffic Images | `GetAllTrafficImages` |
| 2.18 | Traffic Incidents | `GetAllTrafficIncidents` |
| 2.19 | Traffic Speed Bands | `GetAllTrafficSpeedBands`, `GetTrafficSpeedBandsPaginated` |
| 2.20 | VMS / EMAS | `GetAllVMS`, `GetVMSPaginated` |
| 2.21 | Bicycle Parking | `GetBicycleParking` |
| 2.22 | Geospatial Whole Island | `GetGeospatialLayer`, `GetAllGeospatialLayers` |
| 2.23 | Facilities Maintenance | `GetAllFacilityMaintenance` |
| 2.24 | Station Crowd Density Real Time | `GetStationCrowdDensityRealTimeByLine`, `GetAllStationCrowdDensityRealTime` |
| 2.25 | Station Crowd Density Forecast | `GetStationCrowdDensityForecastByLine`, `GetAllStationCrowdDensityForecast` |
| 2.26 | Traffic Flow | `GetTrafficFlow` |
| 2.27 | Planned Bus Routes | `GetAllPlannedBusRoutes`, `GetPlannedBusRoutesPaginated` |
| 2.28 | Electric Vehicle Charging Points | `GetEVChargersByPostalCode` |
| 2.29 | Electric Vehicle Charging Points Batch | `GetEVChargersBatch` |
| 2.30 | Flood Alerts | `GetAllFloodAlerts` |
| 2.31 | GTFS Schedule (Train) | `GetGTFSScheduleTrain` |
| 2.32 | GTFS Realtime (Train Service Alerts) | `GetGTFSRealTimeTrainServiceAlerts` |
| 2.33 | GTFS Realtime (Train Trip Updates - Disruption) | `GetGTFSRealTimeTrainTripUpdates` |
