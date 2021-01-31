package fetcher

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Client interface {
	Get(string) (string, error)
}

type HttpClient struct {
	client http.Client
}

func (c HttpClient) Get(url string) (string, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		log.Println("Error fetching url:", url)
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Could not read body for request:", url)
		return "", err
	}

	return string(body), err
}

func NewClient() Client {
	return HttpClient{http.Client{}}
}
