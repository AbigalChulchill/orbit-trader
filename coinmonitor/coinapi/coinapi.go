package coinapi

const BaseURL = "https://rest.coinapi.io"

type Client struct {
	apiKey string
}

func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}