package fetcher

import (
	"errors"
	"testing"
	"time"
)

var (
	BBCFeedWithoutTTL = RSSFeed{
		Link: "http://old-link.com",
	}
	OriginalBBCFeed = RSSFeed{
		Link: "http://old-link.com",
		TTL:  10,
	}
	FreshBBCFeed = RSSFeed{
		Link: "http://new-link.com",
		TTL:  10,
	}
)

func TestFetchingSecondTimeUsesCachedResponse(t *testing.T) {
	fetcher := fetcherWithStubResponse(bbcURL, OriginalBBCFeed)
	cachedFetcher := NewCachedFetcher(fetcher)

	rssFeed, err := cachedFetcher.Fetch(bbcURL)
	if err != nil {
		t.Fatal("Failed with error:", err)
	}

	// Set new response in fetcher
	cachedFetcher.fetcher = fetcherWithStubResponse(bbcURL, FreshBBCFeed)

	rssFeed, err = cachedFetcher.Fetch(bbcURL)
	if err != nil {
		t.Fatal("Failed with error:", err)
	}

	if rssFeed.Link != OriginalBBCFeed.Link {
		t.Fatalf("Should not have called the fetcher a second time, expected to see the old link instead got: %s", rssFeed.Link)
	}
}

func TestFeedsWithoutATTLGetADefaultTTL(t *testing.T) {
	now := time.Now()
	timeNow := func() time.Time { return now }
	timeBeforeExpiry := func() time.Time { return now.Add(DefaultTTL - time.Second) }
	timeAtExpiry := func() time.Time { return now.Add(DefaultTTL) }

	fetcher := fetcherWithStubResponse(bbcURL, BBCFeedWithoutTTL)
	cachedFetcher := NewCachedFetcher(fetcher)
	// Replace time function with a fixed known moment.
	cachedFetcher.timeNow = timeNow

	rssFeed, err := cachedFetcher.Fetch(bbcURL)
	if err != nil {
		t.Fatal("Failed with error:", err)
	}

	// Set new response in fetcher
	cachedFetcher.fetcher = fetcherWithStubResponse(bbcURL, FreshBBCFeed)

	// Right Before expiring
	cachedFetcher.timeNow = timeBeforeExpiry

	rssFeed, err = cachedFetcher.Fetch(bbcURL)
	if err != nil {
		t.Fatal("Failed with error:", err)
	}

	if rssFeed.Link != OriginalBBCFeed.Link {
		t.Fatalf("Should still be using the old feed, expected to see the old link instead got: %s", rssFeed.Link)
	}

	// After expiring
	// Move time forward to when our cached copy should expire
	cachedFetcher.timeNow = timeAtExpiry

	rssFeed, err = cachedFetcher.Fetch(bbcURL)
	if err != nil {
		t.Fatal("Failed with error:", err)
	}

	if rssFeed.Link == OriginalBBCFeed.Link {
		t.Fatalf("Should have fetched the fresh feed, expected to see the new link instead got: %s", rssFeed.Link)
	}
}

func TestFetchingSecondTimeWithSetTTLAfterExpiryGetsFreshContent(t *testing.T) {
	now := time.Now()
	timeNow := func() time.Time { return now }
	timeAtExpiry := func() time.Time { return now.Add(time.Minute * time.Duration(OriginalBBCFeed.TTL)) }
	timeBeforeExpiry := func() time.Time { return now.Add(DefaultTTL - time.Second) }

	fetcher := fetcherWithStubResponse(bbcURL, OriginalBBCFeed)
	cachedFetcher := NewCachedFetcher(fetcher)
	// Replace time function with a fixed known moment.
	cachedFetcher.timeNow = timeNow

	rssFeed, err := cachedFetcher.Fetch(bbcURL)
	if err != nil {
		t.Fatal("Failed with error:", err)
	}

	// Set new response in fetcher
	cachedFetcher.fetcher = fetcherWithStubResponse(bbcURL, FreshBBCFeed)

	// Right Before expiring
	cachedFetcher.timeNow = timeBeforeExpiry

	rssFeed, err = cachedFetcher.Fetch(bbcURL)
	if err != nil {
		t.Fatal("Failed with error:", err)
	}

	if rssFeed.Link != OriginalBBCFeed.Link {
		t.Fatalf("Should still be using the old feed, expected to see the old link instead got: %s", rssFeed.Link)
	}

	// After expiring
	// Move time forward to when our cached copy should expire
	cachedFetcher.timeNow = timeAtExpiry

	rssFeed, err = cachedFetcher.Fetch(bbcURL)
	if err != nil {
		t.Fatal("Failed with error:", err)
	}

	if rssFeed.Link == OriginalBBCFeed.Link {
		t.Fatalf("Should have fetched the fresh feed, expected to see the new link instead got: %s", rssFeed.Link)
	}
}

func TestShouldReturnTheFetchingErrorIfFailingOnFirstFetch(t *testing.T) {
	expectedError := errors.New("Fetch failed")
	fetcher := fetcherWithErrorResponse(bbcURL, expectedError)
	cachedFetcher := NewCachedFetcher(fetcher)

	_, err := cachedFetcher.Fetch(bbcURL)
	if err == nil {
		t.Fatal("Expected to fail with error", expectedError)
	}
	if err != expectedError {
		t.Fatal("Expected to fail with error: ", expectedError, ". Got", err)
	}
}

func TestFetchingSecondTimeAfterExpiryOnErrorServesExpiredContent(t *testing.T) {
	now := time.Now()
	timeNow := func() time.Time { return now }
	timeAtExpiry := func() time.Time { return now.Add(time.Minute * time.Duration(OriginalBBCFeed.TTL)) }

	fetcher := fetcherWithStubResponse(bbcURL, OriginalBBCFeed)
	cachedFetcher := NewCachedFetcher(fetcher)
	// Replace time function with a fixed known moment.
	cachedFetcher.timeNow = timeNow

	rssFeed, err := cachedFetcher.Fetch(bbcURL)
	if err != nil {
		t.Fatal("Failed with error:", err)
	}

	// Set new response in fetcher
	cachedFetcher.fetcher = fetcherWithErrorResponse(bbcURL, errors.New("Fetch failed"))
	// Move time forward to when our cached copy should expire
	cachedFetcher.timeNow = timeAtExpiry

	rssFeed, err = cachedFetcher.Fetch(bbcURL)
	if err != nil {
		t.Fatal("Failed with error:", err)
	}

	if rssFeed.Link != OriginalBBCFeed.Link {
		t.Fatalf("Should have reused the old cached feed, expected to see the old link instead got: %s", rssFeed.Link)
	}
}
