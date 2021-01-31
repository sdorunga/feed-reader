package fetcher

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

const (
	bbcURL = "http://feeds.bbci.co.uk/news/uk/rss.xml"
)

func TestFetchRSSFeed(t *testing.T) {
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
		Link:        "https://www.bbc.co.uk/news/",
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
		t.Fatalf("\nExpected:\n  %v.\nGot:\n  %v\n", expectedRSSFeed, rssFeed)
	}
}

func TestFetchRSSFeedWithEmptyFieldsFiltersUselessItems(t *testing.T) {
	client := clientWithStubResponse(MissingFieldsBBCResponse)
	fetcher := FeedFetcher{client}

	rssFeed, err := fetcher.Fetch(bbcURL)
	if err != nil {
		t.Fatal("Failed with error:", err)
	}

	expectedRSSFeed := RSSFeed{
		Title:       "BBC News - Home",
		Description: "BBC News - Home",
		ImageURL:    "https://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif",
		Link:        "https://www.bbc.co.uk/news/",
		Items: []FeedItem{
			FeedItem{
				Title:       "EU vaccine export row: Bloc backtracks on controls for NI",
				Description: "It follows a decision to invoke an emergency provision in the Brexit deal in order to control vaccine exports.",
				Link:        "https://www.bbc.co.uk/news/uk-55865539",
			},
		},
	}

	if !reflect.DeepEqual(rssFeed, expectedRSSFeed) {
		t.Fatalf("\nExpected:\n  %v.\nGot:\n  %v\n", expectedRSSFeed, rssFeed)
	}
}
func TestFetchRSSFeedWithWrongVersion(t *testing.T) {
	client := clientWithStubResponse(OldVersionBBCResponse)
	fetcher := FeedFetcher{client}

	_, err := fetcher.Fetch(bbcURL)
	if err == nil {
		t.Fatal("Expected to fail parsing because of a wrong RSS version")
	}

	if err != ErrorIncompatibleRSSVersion {
		t.Fatalf("\nExpected to fail because of RSS version.\nInstead failed with:\n  %v\n", err)
	}
}

func TestFetchRSSFeedWithMissingRequiredChannelField(t *testing.T) {
	client := clientWithStubResponse(InvalidChannelBBCResponse)
	fetcher := FeedFetcher{client}

	_, err := fetcher.Fetch(bbcURL)
	if err == nil {
		t.Fatal("Expected to fail parsing because of missing required channel fields")
	}

	if err != ErrorRSSInvalidFormat {
		t.Fatalf("\nExpected to fail because of missing required channel field.\nInstead failed with:\n  %v\n", err)
	}
}

func TestFetchRSSFeedWithInvalidXML(t *testing.T) {
	client := clientWithStubResponse(InvalidXMLBBCResponse)
	fetcher := FeedFetcher{client}

	_, err := fetcher.Fetch(bbcURL)
	if err == nil {
		t.Fatal("Expected to fail parsing because of invalid XML")
	}

	if err != ErrorUnparseableXML {
		t.Fatalf("\nExpected to fail because of invalid XML.\nInstead failed with:\n  %v\n", err)
	}
}

func TestFetchRSSFeedWithEmptyXML(t *testing.T) {
	client := clientWithStubResponse(EmptyXMLBBCResponse)
	fetcher := FeedFetcher{client}

	_, err := fetcher.Fetch(bbcURL)
	if err == nil {
		t.Fatal("Expected to fail parsing because of empty XML")
	}

	if err != ErrorRSSInvalidFormat {
		t.Fatalf("\nExpected to fail because of empty RSS.\nInstead failed with:\n  %v\n", err)
	}
}

func TestFetchRSSFeedReturnsNetworkErrorWhenNetworkFails(t *testing.T) {
	client := clientWithErrorResponse(errors.New("No connection"))
	fetcher := FeedFetcher{client}

	_, err := fetcher.Fetch(bbcURL)
	if err == nil {
		t.Fatal("Expected to fail because of network error")
	}

	if err != ErrorNetworkError {
		t.Fatalf("\nExpected to fail because of network error.\nInstead failed with:\n  %v\n", err)
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
