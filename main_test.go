package ltadatamall

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

var testClient *APIClient
var currentYear int
var currentMonth time.Month

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

	currentYear, currentMonth, _ = time.Now().Date()
	if currentMonth <= 2 {
		currentYear = currentYear - 1
		currentMonth = 11 - 1 + currentMonth
	}

	testClient = NewClient(baseUrl, apiKey)
}
