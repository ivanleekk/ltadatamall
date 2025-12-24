package ltadatamall

import (
	"testing"
)

func TestGetAllTaxiAvailability(t *testing.T) {
	response, err := GetAllTaxiAvailability(testClient)
	if err != nil {
		t.Fatalf("Error calling GetTaxiAvailability: %v", err)
	}

	if len(response.Taxis) <= 500 {
		t.Errorf("Expected more than 500 Taxis in response")
	}
}

func TestGetTaxiAvailabilityPaginated(t *testing.T) {
	response, err := GetTaxiAvailabilityPaginated(testClient, 500)
	if err != nil {
		t.Fatalf("Error calling GetTaxiAvailabilityPaginated: %v", err)
	}
	if len(response.Taxis) == 0 {
		t.Errorf("Expected non-empty Taxis in response")
	}
}

func TestGetTaxiAvailabilityPaginated_NoMoreData(t *testing.T) {
	_, err := GetTaxiAvailabilityPaginated(testClient, 1000000)
	if err == nil {
		t.Fatalf("Expected error when no more taxis are available")
	}
}

// TaxiStands API gets rate-limited too quickly, therefore we are not testing it as it will fail
//func TestGetAllTaxiStands(t *testing.T) {
//	response, err := GetAllTaxiStands(testClient)
//	if err != nil {
//		t.Fatalf("Error calling GetAllTaxiAvailability: %v", err)
//	}
//
//	if len(response.TaxiStands) <= 500 {
//		t.Errorf("Expected more than 500 Taxi Stands in response")
//	}
//}
//
//func TestGetTaxiStandsPaginated(t *testing.T) {
//	response, err := GetTaxiStandsPaginated(testClient, 500)
//	if err != nil {
//		t.Fatalf("Error calling GetTaxiStandsPaginated: %v", err)
//	}
//	if len(response.TaxiStands) == 0 {
//		t.Errorf("Expected non-empty Taxis in response")
//	}
//	log.Default().Println(len(response.TaxiStands))
//}
//
//func TestGetTaxiStandsPaginated_NoMoreData(t *testing.T) {
//	_, err := GetTaxiStandsPaginated(testClient, 1000000)
//	if err == nil {
//		t.Fatalf("Expected error when no more taxis are available")
//	}
//}
