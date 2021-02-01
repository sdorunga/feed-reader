package app

import (
	"feed-reader/feedlist"
	"feed-reader/fetcher"
	"time"
)

type Feed struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

type RSSFeed struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Items       []Item `json:"items"`
}

type Item struct {
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	Link           string     `json:"link"`
	PublishingDate *time.Time `json:"pubDate,omitempty"`
	GUID           string     `json:"guid,omitempty"`
	Categories     []string   `json:"categories,omitempty"`
}

func ToResponseFeed(domainFeed feedlist.Feed) Feed {
	return Feed{
		ID:          domainFeed.ID,
		Title:       domainFeed.Title,
		Link:        domainFeed.URL,
		Description: domainFeed.Description,
	}
}

func ToResponseFeeds(domainFeeds []feedlist.Feed) []Feed {
	result := []Feed{}
	for _, domainFeed := range domainFeeds {
		result = append(result, ToResponseFeed(domainFeed))
	}
	return result
}

// Note: normally the pagination would not be here and I would pass it through
// to the source of the feed items.  In this case we always receive the full
// feed so it doesn't provide the normal performance benefits.
func ToResponseRSSFeed(ID string, domainFeed fetcher.RSSFeed, category string, perPage, offset int) RSSFeed {
	items := domainFeed.Items
	if category != "" {
		items = domainFeed.ItemsByCategory(category)
	}
	if perPage != 0 {
		start := offset
		end := offset + perPage

		if start > len(items) {
			start = len(items)
		}
		if end > len(items) {
			end = len(items)
		}

		items = items[start:end]
	}
	return RSSFeed{
		ID:          ID,
		Title:       domainFeed.Title,
		Link:        domainFeed.Link,
		Description: domainFeed.Description,
		ImageURL:    domainFeed.ImageURL,
		Items:       ToResponseRSSItems(items),
	}
}

func ToResponseRSSItems(domainItems []fetcher.FeedItem) []Item {
	items := []Item{}
	for _, domainItem := range domainItems {
		items = append(items, ToResponseRSSItem(domainItem))
	}
	return items
}

func ToResponseRSSItem(domainItem fetcher.FeedItem) Item {
	item := Item{
		Title:       domainItem.Title,
		Description: domainItem.Description,
		Link:        domainItem.Link,
		Categories:  domainItem.Categories,
	}
	if item.PublishingDate != (&time.Time{}) {
		item.PublishingDate = &domainItem.PublishingDate.Time
	}
	return item
}
