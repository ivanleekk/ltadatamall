package ltadatamall

import (
	"net/http"
)

type APIClient struct {
	baseURL    string
	accountKey string
	httpClient *http.Client
}

func NewClient(baseURL, accountKey string) *APIClient {
	return &APIClient{
		baseURL:    baseURL,
		accountKey: accountKey,
		httpClient: &http.Client{},
	}
}
