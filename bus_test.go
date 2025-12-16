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

func TestGetAllBusService(t *testing.T) {
	baseURL := os.Getenv("LTA_DATAMALL_BASE_URL")
	accountKey := os.Getenv("LTA_DATAMALL_ACCOUNT_KEY")
	if baseURL == "" || accountKey == "" {
		t.Fatalf("Environment variables LTA_DATAMALL_BASE_URL and LTA_DATAMALL_ACCOUNT_KEY must be set")
	}

	client := NewClient(baseURL, accountKey)

	response, err := GetAllBusService(client)

	if err != nil {
		t.Fatalf("Error calling GetAllBusService: %v", err)
	}

	if len(response.BusServices) == 0 {
		t.Errorf("Expected non-empty BusServices in response")
	}
}

func TestGetAllBusRoutes(t *testing.T) {
	baseURL := os.Getenv("LTA_DATAMALL_BASE_URL")
	accountKey := os.Getenv("LTA_DATAMALL_ACCOUNT_KEY")
	if baseURL == "" || accountKey == "" {
		t.Fatalf("Environment variables LTA_DATAMALL_BASE_URL and LTA_DATAMALL_ACCOUNT_KEY must be set")
	}

	client := NewClient(baseURL, accountKey)

	response, err := GetAllBusRoute(client)

	if err != nil {
		t.Fatalf("Error calling GetAllBusRoute: %v", err)
	}

	if len(response.BusRoutes) == 0 {
		t.Errorf("Expected non-empty BusRoutes in response")
	}
}
