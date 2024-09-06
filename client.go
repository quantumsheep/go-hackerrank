package hackerrank

import "net/http"

type Client struct {
	apiKey string

	httpClient *http.Client
}

type ClientOptions func(*Client)

func NewClient(apiKey string, options ...ClientOptions) *Client {
	client := &Client{
		apiKey: apiKey,

		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func WithHttpClient(httpClient *http.Client) ClientOptions {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}
