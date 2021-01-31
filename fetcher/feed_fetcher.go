package fetcher

import (
	"encoding/xml"
	"errors"
	"log"
	"time"
)

var (
	// ErrorIncompatibleRSSVersion indicates the RSS feed version is different
	// to 2.0.  The fetcher only supports RSS 2.0 for simplicty
	ErrorIncompatibleRSSVersion = errors.New("Incompatible RSS version")
	// ErrorRSSInvalidFormat is returned when the required fields for an RSS
	// feed are not present.
	ErrorRSSInvalidFormat = errors.New("Invalid RSS Format")
	// ErrorNetworkError is returned whenever the underlying client fails to
	// make a network request
	ErrorNetworkError = errors.New("Client failed to fetch feed")
	// ErrorUnparseableXML is returned when the XML feed is not correctly
	// formatted and is unparseable by the GO XML parser.
	ErrorUnparseableXML = errors.New("XML is not valid")
)

func NewFeedFetcher() FeedFetcher {
	return FeedFetcher{NewClient()}
}

// FeedFetcher is responsible for downloading an RSS feed and parsing it into
// an internal representation.  It will only fetch feeds that are RSS 2.0
// compatible.
type FeedFetcher struct {
	client Client
}

// RSSFeed is an internal version of the RSS xml format. It mostly matches what
// is in the specification but with a few small tweaks to make it more
// convenient.
type RSSFeed struct {
	Title       string     `xml:"title"`
	Description string     `xml:"description"`
	ImageURL    string     `xml:"image>url"`
	Link        string     `xml:"link"`
	Items       []FeedItem `xml:"item"`
	TTL         int        `xml:"ttl"`
}

// FeedItem is an individual news item.
type FeedItem struct {
	Title          string      `xml:"title"`
	Description    string      `xml:"description"`
	Link           string      `xml:"link"`
	PublishingDate RFC1132Time `xml:"pubDate"`
	GUID           string      `xml:"guid"`
}

// rssDoc is just a wrapper struct that we use to parse and validate the XML
// document. It's not actually used interally in the application.
type rssDoc struct {
	Version string  `xml:"version,attr"`
	Channel RSSFeed `xml:"channel"`
}

func (doc *rssDoc) Validate() error {
	// Note: I should be more flexible here in real life and actually parse the
	// version and check for compatiblity. Given that RSS is a pretty fixed
	// standard and has been on a stable version since 2009 and all the example
	// feeds are on 2.0 this is fine for the exercise.
	if doc.Version != "2.0" {
		return ErrorIncompatibleRSSVersion
	}

	switch {
	case doc.Channel.Title == "",
		doc.Channel.Description == "",
		doc.Channel.Link == "":
		return ErrorRSSInvalidFormat
	}

	doc.CleanItems()

	return nil
}

// Note: These fields are not srictly required by the RSS standard but
// without them we can't meaningfully display the item in a feed
// so instead of failing I just clear them away
func (doc *rssDoc) CleanItems() {
	cleanedItems := []FeedItem{}

	for _, item := range doc.Channel.Items {
		if item.Title != "" &&
			item.Description != "" &&
			item.Link != "" {
			cleanedItems = append(cleanedItems, item)
		}
	}

	doc.Channel.Items = cleanedItems
}

// RFC1132Time is a custom type so that we can have a custom decoder to
// unmarshal the time string we get from RSS.  This is the time format that is
// in the RSS spec.
type RFC1132Time struct {
	time.Time
}

// UnmarshalXML should only be called by the Go XML unmarshaller, it's only
// public so it can fulfil the interface
func (c *RFC1132Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parsedTime, err := time.Parse(time.RFC1123, v)
	if err != nil {
		log.Printf("Time %s in RSS feed is unparseable: %v", v, err)
		return err
	}

	*c = RFC1132Time{parsedTime}
	return nil
}

// Fetch downloads and parses the RSS feed of one url. It can error out based on the version of RSS
// or because of various network and connection errors.
// Note:
// This could use an XML Decoder to parse the stream from the feed as it comes
// in. I built it this way because it makes the code more straightforward and
// the performance benefits in this particular case are not significant.
func (fetcher FeedFetcher) Fetch(url string) (RSSFeed, error) {
	xmlFeed, err := fetcher.client.Get(url)
	if err != nil {
		log.Printf("Got an error fetching the RSS feed for %s: %v", url, err)
		return RSSFeed{}, ErrorNetworkError
	}

	rssDoc := rssDoc{}
	err = xml.Unmarshal([]byte(xmlFeed), &rssDoc)
	if err != nil {
		log.Printf("XML feed was unparseable: %v", err)
		return RSSFeed{}, ErrorUnparseableXML
	}

	err = rssDoc.Validate()
	if err != nil {
		return RSSFeed{}, err
	}

	return rssDoc.Channel, nil
}
