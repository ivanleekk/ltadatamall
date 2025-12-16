package ltadatamall

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env from the project root
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var testClient *APIClient

func TestMain(m *testing.M) {
	baseURL := os.Getenv("LTA_DATAMALL_BASE_URL")
	accountKey := os.Getenv("LTA_DATAMALL_ACCOUNT_KEY")
	if baseURL == "" || accountKey == "" {
		os.Exit(1)
	}

	testClient = NewClient(baseURL, accountKey)

	os.Exit(m.Run())
}

func TestGetAllBusServices(t *testing.T) {
	response, err := GetAllBusServices(testClient)
	if err != nil {
		t.Fatalf("Error calling GetAllBusService: %v", err)
	}

	if len(response.BusServices) == 0 {
		t.Errorf("Expected non-empty BusServices in response")
	}
}

func TestGetAllBusRoutes(t *testing.T) {
	response, err := GetAllBusRoutes(testClient)

	if err != nil {
		t.Fatalf("Error calling GetAllBusRoute: %v", err)
	}

	if len(response.BusRoutes) == 0 {
		t.Errorf("Expected non-empty BusRoutes in response")
	}
}

func TestGetAllBusStops(t *testing.T) {
	response, err := GetAllBusStops(testClient)

	if err != nil {
		t.Fatalf("Error calling GetAllBusStops: %v", err)
	}

	if len(response.BusStops) == 0 {
		t.Errorf("Expected non-empty BusStops in response")
	}
}
