package app

import (
	"feed-reader/feedlist"
	"feed-reader/fetcher"

	"net/url"
)

type GETFeedHandler struct {
	store   feedlist.FeedListStore
	fetcher fetcher.Fetcher
}

type GETFeedResponse struct {
	Feed RSSFeed `json:"feed"`
}

func (handler GETFeedHandler) Handle(body []byte, params map[string]string, queryParams url.Values) (interface{}, error) {
	feed, err := handler.store.GetByID(params["id"])
	if err != nil {
		if err == feedlist.ErrorFeedNotFound {
			return nil, NotFoundError{err: err}
		}
		return nil, err
	}

	rssFeed, err := handler.fetcher.Fetch(feed.URL)
	if err != nil {
		return nil, err
	}

	return GETFeedResponse{
		Feed: ToResponseRSSFeed(feed.ID, rssFeed, queryParams.Get("category")),
	}, nil
}
