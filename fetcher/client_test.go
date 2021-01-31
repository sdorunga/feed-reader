package fetcher

type TestClient struct {
	stubResponse string
	err          error
}

func (client TestClient) Get(url string) (string, error) {
	if client.err != nil {
		return "", client.err
	}

	return client.stubResponse, nil
}

func clientWithStubResponse(stubResponse string) Client {
	return TestClient{stubResponse, nil}
}

func clientWithErrorResponse(err error) Client {
	return TestClient{"", err}
}
