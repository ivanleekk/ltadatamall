package ltadatamall

type APIClient struct {
	baseURL    string
	accountKey string
}

func NewClient(baseURL, accountKey string) *APIClient {
	return &APIClient{
		baseURL:    baseURL,
		accountKey: accountKey,
	}
}
