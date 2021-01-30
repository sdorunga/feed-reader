package fetcher

type TestClient struct {
	stubResponse string
}

func (client TestClient) Get(url string) (string, error) {
	return client.stubResponse, nil
}

func clientWithStubResponse(stubResponse string) Client {
	return TestClient{stubResponse}
}
