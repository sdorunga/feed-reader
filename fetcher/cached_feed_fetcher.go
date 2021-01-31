package fetcher

import (
	"log"
	"sync"
	"time"
)

const (
	// DefaultTTL is how long we'll cache a feed that doesn't have a TTL
	// explictly set in the response we get from the provider. 1 Minute seems
	// like a good compromise btween freshness and not hitting the 3rd party
	// very often.
	DefaultTTL = 1 * time.Minute
)

// Fetcher is the interface that a Feed fetcher has to fulfil. Both the Cached and normal
// version implement this
type Fetcher interface {
	Fetch(string) (RSSFeed, error)
}

// CachedFeedFetcher is a wrapper around the FeedFetcher that caches responses based on the
// TTL that comes back from the feed. If not present it falls back to DefaultTTL
//
// The implementation is thread-safe so it can be used across requests
type CachedFeedFetcher struct {
	fetcher Fetcher
	// Note: The cache is potentially unbounded. Given the size of a typical
	// RSS feed and the relatively limited amount of feeds that the app will
	// serve this is ok.
	cache   map[string]cachedFeed
	timeNow func() time.Time
	mu      *sync.RWMutex
}

// cachedFeed is an expirable entry in the cache
type cachedFeed struct {
	feed      RSSFeed
	expiresAt time.Time
}

// NewCachedFetcher returns a fully initialized version of a CachedFeedFetcher
// and is the only way that one should be created
func NewCachedFetcher(fetcher Fetcher) CachedFeedFetcher {
	return CachedFeedFetcher{
		fetcher: fetcher,
		cache:   map[string]cachedFeed{},
		timeNow: time.Now,
		mu:      &sync.RWMutex{},
	}
}

func (cachedFetcher CachedFeedFetcher) Fetch(url string) (RSSFeed, error) {
	cachedFeedEntry, entryIsCached := cachedFetcher.getCachedEntry(url)
	if entryIsCached {
		if cachedFetcher.timeNow().Before(cachedFeedEntry.expiresAt) {
			return cachedFeedEntry.feed, nil
		}
	}

	// Note: We could lock for writing here and have each locker check again if
	// the cache is now up to date. This would ensure we only make one request
	// at every expiry. I chose to not do this because during this request we
	// would be blocking reads and I made the assumption that a slightly out of
	// date feed is ok from a user perspective.
	feed, err := cachedFetcher.fetcher.Fetch(url)
	if err != nil {
		// Note: When we fail to fetch and have a cached version we return
		// the cached version even if it is possibly out of date In this
		// case I think it's preferable to have a stale feed rather than
		// not return anything at all.
		if entryIsCached {
			log.Printf("Failed to fetch %s, falling back to cached version, it is %d seconds old", url, cachedFetcher.timeNow().Sub(cachedFeedEntry.expiresAt))
			return cachedFeedEntry.feed, nil
		}
		return RSSFeed{}, err
	}

	cachedFetcher.putCachedEntry(url, feed)

	return feed, nil
}

func (cachedFetcher CachedFeedFetcher) getCachedEntry(url string) (cachedFeed, bool) {
	cachedFetcher.mu.RLock()
	defer cachedFetcher.mu.RUnlock()

	cachedEntry, entryIsCached := cachedFetcher.cache[url]
	return cachedEntry, entryIsCached
}

func (cachedFetcher CachedFeedFetcher) putCachedEntry(url string, feed RSSFeed) {
	expiresAt := cachedFetcher.timeNow().Add(DefaultTTL)
	if feed.TTL != 0 {
		expiresAt = cachedFetcher.timeNow().Add(time.Duration(feed.TTL) * time.Minute)
	}

	cachedFetcher.mu.Lock()
	defer cachedFetcher.mu.Unlock()

	cachedFetcher.cache[url] = cachedFeed{
		feed:      feed,
		expiresAt: expiresAt,
	}
}
