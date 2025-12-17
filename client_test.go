package ltadatamall

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	baseURL := "https://datamall2.mytransport.sg/ltaodataservice/"
	accountKey := "test_account_key"

	client := NewClient(baseURL, accountKey)

	if client.baseURL != baseURL {
		t.Errorf("Expected baseURL %s, got %s", baseURL, client.baseURL)
	}

	if client.accountKey != accountKey {
		t.Errorf("Expected accountKey %s, got %s", accountKey, client.accountKey)
	}

	if client.httpClient == nil {
		t.Error("Expected httpClient to be initialized, got nil")
	}
}
