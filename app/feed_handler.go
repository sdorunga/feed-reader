package app

import (
	"feed-reader/feedlist"
	"feed-reader/fetcher"

	"net/url"
)

// GETFeedHandler handles requests to `GET /feeds/<id>` with the following optional params:
//   - `category`, limited to 1 allows for filtering on the category field of
//   the Feed items if present
//   - `per_page` number of feed items to display per page
//   - `offset` number of feed items to skip
//
// Returns a GETFeedResponse which has a list of articles in the RSS feed
type GETFeedHandler struct {
	store   feedlist.FeedListStore
	fetcher fetcher.Fetcher
}

// GETFeedResponse represents the format of the response for this endpoint
// `
// {
//     "feed": {
//         "description": "<Short description of the content of the feed>",
//         "id": "<UUID of the feed in the application database>",
//         "image_url": "<URL for an image of the provider>",
//         "items": [
//             {
//                 "description": "<Short description of the article>",
//                 "link": "<URL to article>",
//                 "pubDate": "<Date of publishing in 2021-01-31T17:05:00Z format>",
//                 "title": "<Title of the article>"
//             },
//         ],
//         "link": "<Link to the RSS feed provider homepage>",
//         "title": "<Title of the feed>"
//     }
// }
// `
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

	perPage, offset, err := getPaginationParams(queryParams)
	if err != nil {
		return nil, err
	}

	return GETFeedResponse{
		Feed: ToResponseRSSFeed(feed.ID, rssFeed, queryParams.Get("category"), perPage, offset),
	}, nil
}
