package fetcher

import (
	"encoding/xml"
	"errors"
	"time"
)

var (
	// ErrorIncompatibleRSSVersion indicates the RSS feed version is different
	// to 2.0.  The fetcher only supports RSS 2.0 for simplicty
	ErrorIncompatibleRSSVersion = errors.New("Incompatible RSS version")
)

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
	Items       []FeedItem `xml:"item"`
}

// FeedItem is an individual news item.
type FeedItem struct {
	Title          string      `xml:"title"`
	Description    string      `xml:"description"`
	Link           string      `xml:"link"`
	TTL            int         `xml:"ttl"`
	PublishingDate RFC1132Time `xml:"pubDate"`
	GUID           string      `xml:"guid"`
}

// rssDoc is just a wrapper struct that we use to parse and validate the XML
// document. It's not actually used interally in the application.
type rssDoc struct {
	Version string  `xml:"version,attr"`
	Channel RSSFeed `xml:"channel"`
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
func (fetcher *FeedFetcher) Fetch(url string) (RSSFeed, error) {
	xmlFeed, err := fetcher.client.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	rssDoc := rssDoc{}
	err = xml.Unmarshal([]byte(xmlFeed), &rssDoc)
	if err != nil {
		return RSSFeed{}, err
	}

	return rssDoc.Channel, nil
}
