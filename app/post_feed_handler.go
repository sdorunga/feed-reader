package app

import (
	"feed-reader/feedlist"
	"feed-reader/fetcher"

	"encoding/json"
	"errors"
	"net/url"
)

// POSTFeedHandler handles requests to `POST /feeds/<id>` with the following required body:
// {
//     "url": "<URL to an RSS feed>"
// }
//
// As well as these optional params:
//   - `per_page` number of feed items to display per page
//   - `offset` number of feed items to skip
//
// It will add the Feed url to the list of supported RSS feeds and retreive it synchronously
// for display in the Client. This also ensures that the Feed is valid.
//
// Returns a PostFeedResponse which has a list of articles in the RSS feed
type POSTFeedHandler struct {
	store   feedlist.FeedListStore
	fetcher fetcher.Fetcher
}

// POSTFeedResponse represents the format of the response for this endpoint.
// It is the same as the one we can later GET from the app. This is to save
// on a round-trip.
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
type POSTFeedResponse struct {
	Feed RSSFeed `json:"feed"`
}

type POSTFeedRequest struct {
	URL string `json:"url"`
}

func (handler POSTFeedHandler) Handle(body []byte, params map[string]string, queryParams url.Values) (interface{}, error) {
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
		ImageURL:    rssFeed.ImageURL,
	})
	if err != nil {
		return nil, err
	}

	perPage, offset, err := getPaginationParams(queryParams)
	if err != nil {
		return nil, err
	}

	return GETFeedResponse{
		Feed: ToResponseRSSFeed(feedID, rssFeed, "", perPage, offset),
	}, nil
}
