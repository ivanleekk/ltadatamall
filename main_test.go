package ltadatamall

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var testClient *APIClient

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	os.Exit(code)
}

func setup() {
	// Load .env from the project root
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Initialize the client once
	apiKey := os.Getenv("LTA_DATAMALL_ACCOUNT_KEY")
	if apiKey == "" {
		log.Fatal("LTA_DATAMALL_ACCOUNT_KEY not set in environment")
	}

	baseUrl := os.Getenv("LTA_DATAMALL_BASE_URL")
	if baseUrl == "" {
		log.Fatal("LTA_DATAMALL_BASE_URL not set in environment")
	}

	testClient = NewClient(baseUrl, apiKey)
}
