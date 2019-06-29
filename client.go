package client

// Client is a Logz API client, it holds the API token
type Client struct {
	apiToken string
}

// New returns a Client using the passed in API token
func New(apiToken string) *Client {
	return &Client{
		apiToken: apiToken,
	}
}
