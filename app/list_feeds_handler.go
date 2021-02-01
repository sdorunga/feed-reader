package app

import (
	"feed-reader/feedlist"

	"net/url"
)

// GETFeedListHandler handles requests to `GET /feeds`:
//
// Returns a GETFeedListResponse which has a list of the currently supported RSS feeds
type GETFeedListHandler struct {
	store feedlist.FeedListStore
}

// GETFeedListResponse represents the format of the response for this endpoint
// `
// {
//     "feed": {
//         "description": "<Short description of the content of the feed>",
//         "id": "<UUID of the feed in the application database>",
//         "image_url": "<URL for an image of the provider>",
//         "link": "<Link to the RSS feed provider homepage>",
//         "title": "<Title of the feed>"
//     }
// }
// `
type GETFeedListResponse struct {
	Feeds []Feed `json:"feeds"`
}

func (handler GETFeedListHandler) Handle(body []byte, params map[string]string, _ url.Values) (interface{}, error) {
	feeds, err := handler.store.ListAll()
	if err != nil {
		return nil, err
	}
	return GETFeedListResponse{
		Feeds: ToResponseFeeds(feeds),
	}, nil
}
