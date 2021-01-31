package app

import (
	"feed-reader/feedlist"
	"feed-reader/fetcher"

	"encoding/json"
	"errors"
	"net/url"
)

type POSTFeedHandler struct {
	store   feedlist.FeedListStore
	fetcher fetcher.Fetcher
}

type POSTFeedResponse struct {
	Feed RSSFeed `json:"feed"`
}

type POSTFeedRequest struct {
	URL string `json:"url"`
}

func (handler POSTFeedHandler) Handle(body []byte, params map[string]string) (interface{}, error) {
	request := POSTFeedRequest{}
	err := json.Unmarshal(body, &request)
	if err != nil {
		return nil, err
	}
	if request.URL == "" {
		return nil, BadRequestError{err: errors.New("Invalid request, URL is required")}
	}
	// Using this as a simple validator for the URL
	url, err := url.Parse(request.URL)
	if err != nil {
		return nil, BadRequestError{err: err}
	}

	rssFeed, err := handler.fetcher.Fetch(url.String())
	if err != nil {
		return nil, err
	}

	feedID, err := handler.store.Add(feedlist.Feed{
		Title:       rssFeed.Title,
		Description: rssFeed.Description,
		URL:         request.URL,
	})
	if err != nil {
		return nil, err
	}

	return GETFeedResponse{
		Feed: ToResponseRSSFeed(feedID, rssFeed),
	}, nil
}
