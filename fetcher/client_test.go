package fetcher

import (
	"errors"
)

type TestClient struct {
	stubResponses map[string]string
	errors        map[string]error
}

var ErrorNoResponseRegistered = errors.New("No registered response for this url")

func (client TestClient) Get(url string) (string, error) {
	if client.errors[url] != nil {
		return "", client.errors[url]
	}

	if response := client.stubResponses[url]; response != "" {
		return response, nil
	}

	return "", ErrorNoResponseRegistered
}

func clientWithStubResponse(url, stubResponse string) Client {
	return TestClient{map[string]string{
		url: stubResponse,
	}, map[string]error{}}
}

func clientWithErrorResponse(url string, err error) Client {
	return TestClient{map[string]string{}, map[string]error{
		url: err,
	}}
}
