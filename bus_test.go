package ltadatamall

import (
	"testing"
)

func TestGetAllBusServices(t *testing.T) {
	response, err := GetAllBusServices(testClient)
	if err != nil {
		t.Fatalf("Error calling GetAllBusService: %v", err)
	}

	if len(response.BusServices) <= 500 {
		t.Errorf("Expected more than 500 BusServices in response")
	}
}

func TestGetBusServicesPaginated(t *testing.T) {
	response, err := GetBusServicesPaginated(testClient, 500)
	if err != nil {
		t.Fatalf("Error calling GetBusServicesPaginated: %v", err)
	}
	if len(response.BusServices) == 0 {
		t.Errorf("Expected non-empty BusServices in response")
	}
}

func TestGetBusServicesPaginated_NoMoreData(t *testing.T) {
	_, err := GetBusServicesPaginated(testClient, 1000000)
	if err == nil {
		t.Fatalf("Expected error when no more bus services are available")
	}
}

func TestGetAllBusRoutes(t *testing.T) {
	response, err := GetAllBusRoutes(testClient)

	if err != nil {
		t.Fatalf("Error calling GetAllBusRoute: %v", err)
	}

	if len(response.BusRoutes) <= 100 {
		t.Errorf("Expected more than 100 BusRoutes in response")
	}
}

func TestGetBusRoutesPaginated(t *testing.T) {
	response, err := GetBusRoutesPaginated(testClient, 500)
	if err != nil {
		t.Fatalf("Error calling GetBusServicesPaginated: %v", err)
	}
	if len(response.BusRoutes) == 0 {
		t.Errorf("Expected non-empty BusServices in response")
	}
}

func TestGetBusRoutesPaginated_NoMoreData(t *testing.T) {
	_, err := GetBusRoutesPaginated(testClient, 1000000)
	if err == nil {
		t.Fatalf("Expected error when no more bus services are available")
	}
}

func TestGetAllBusStops(t *testing.T) {
	response, err := GetAllBusStops(testClient)

	if err != nil {
		t.Fatalf("Error calling GetAllBusRoute: %v", err)
	}

	if len(response.BusStops) <= 500 {
		t.Errorf("Expected more than 500 BusStops in response")
	}
}

func TestGetBusStopsPaginated(t *testing.T) {
	response, err := GetBusStopsPaginated(testClient, 500)
	if err != nil {
		t.Fatalf("Error calling GetBusStopsPaginated: %v", err)
	}
	if len(response.BusStops) == 0 {
		t.Errorf("Expected non-empty BusStops in response")
	}
}

func TestGetBusStopsPaginated_NoMoreData(t *testing.T) {
	_, err := GetBusStopsPaginated(testClient, 1000000)
	if err == nil {
		t.Fatalf("Expected error when no more bus stops are available")
	}
}

func TestGetBusArrivalAtBusStop(t *testing.T) {
	busStopCode := "77009"
	response, err := GetBusArrivalAtBusStop(testClient, busStopCode)

	if err != nil {
		t.Fatalf("Error calling GetBusArrivalAtBusStop: %v", err)
	}

	if len(response.Services) == 0 {
		t.Errorf("Expected non-empty Services in response for bus stop %s", busStopCode)
	}
}

func TestGetBusArrivalAtBusStopAndService(t *testing.T) {
	busStopCode := "77009"
	service := "17"
	response, err := GetBusArrivalAtBusStopAndService(testClient, busStopCode, service)

	if err != nil {
		t.Fatalf("Error calling GetBusArrivalAtBusStop: %v", err)
	}
	if len(response.Services) == 0 {
		t.Errorf("Expected non-empty Services in response for bus stop %s and service %s", busStopCode, service)
	}
}
