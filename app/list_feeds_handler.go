package app

import (
	"feed-reader/feedlist"
)

type GETFeedListHandler struct {
	store feedlist.FeedListStore
}

type GETFeedListResponse struct {
	Feeds []Feed `json:"feeds"`
}

func (handler GETFeedListHandler) Handle(body []byte, params map[string]string) (interface{}, error) {
	feeds, err := handler.store.ListAll()
	if err != nil {
		return nil, err
	}
	return GETFeedListResponse{
		Feeds: ToResponseFeeds(feeds),
	}, nil
}