package fetcher

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

const (
	bbcURL = "http://feeds.bbci.co.uk/news/uk/rss.xml"
)

func TestFetchFreshRSSFeed(t *testing.T) {
	client := clientWithStubResponse(GoodBBCResponse)
	fetcher := FeedFetcher{client}

	rssFeed, err := fetcher.Fetch(bbcURL)
	if err != nil {
		t.Error("Failed with error:", err)
	}

	expectedRSSFeed := RSSFeed{
		Title:       "BBC News - Home",
		Description: "BBC News - Home",
		ImageURL:    "https://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif",
		Items: []FeedItem{
			FeedItem{
				Title:          "EU vaccine export row: Bloc backtracks on controls for NI",
				Description:    "It follows a decision to invoke an emergency provision in the Brexit deal in order to control vaccine exports.",
				Link:           "https://www.bbc.co.uk/news/uk-55865539",
				PublishingDate: forceParseTime("Sat, 30 Jan 2021 14:15:18 GMT"),
				GUID:           "https://www.bbc.co.uk/news/uk-55865539",
			},
			FeedItem{
				Title:          "Arlene Foster urges PM to replace 'unworkable' NI Brexit deal",
				Description:    "Arlene Foster wants GB-NI trade flow problems addressed, in the wake of the EU vaccine export row.",
				Link:           "https://www.bbc.co.uk/news/uk-northern-ireland-55866285",
				PublishingDate: forceParseTime("Sat, 30 Jan 2021 14:21:14 GMT"),
				GUID:           "https://www.bbc.co.uk/news/uk-northern-ireland-55866285",
			},
		},
	}

	if !reflect.DeepEqual(rssFeed, expectedRSSFeed) {
		t.Errorf("\nExpected:\n  %v.\nGot:\n  %v\n", expectedRSSFeed, rssFeed)
	}
}

// forceParseTime is a convenience method to allow us to ignore the possible
// error cases. It should only be used to parse test time strings when they are
// hardcoded and guaranteed to parse.
func forceParseTime(timeString string) RFC1132Time {
	parsedTime, err := time.Parse(time.RFC1123, timeString)
	if err != nil {
		panic(fmt.Sprintf("Unparseable time string: %v!", timeString))
	}

	return RFC1132Time{parsedTime}
}
