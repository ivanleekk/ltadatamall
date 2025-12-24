// ltadatamall package provides a wrapper around Singapore's LTA Datamall APIs
// allowing for easier usage in your own projects with your own API Key
package ltadatamall

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (apiClient *APIClient) get(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, apiClient.baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("AccountKey", apiClient.accountKey)

	resp, err := apiClient.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (apiClient *APIClient) getJSON(endpoint string, out any) error {
	if out == nil {
		return errors.New("out must not be nil")
	}

	resp, err := apiClient.get(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("GET %s: unexpected status code %d", endpoint, resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return err
	}

	return nil
}
