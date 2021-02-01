package feedlist

import (
	"github.com/boltdb/bolt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestHasTheDefaultFeedListWhenEmpty(t *testing.T) {
	db := newTestDB()
	list, err := NewFeedListStore(db)
	if err != nil {
		t.Fatal("Expected no error")
	}

	feedList, err := list.ListAll()
	if err != nil {
		t.Fatal("Expected no error")
	}

	if !reflect.DeepEqual(feedList, defaultFeedsList) {
		t.Fatal("Expected to have the default list of feeds on empty DB")
	}
}

func TestCantAddAnEntryTwice(t *testing.T) {
	db := newTestDB()
	list, err := NewFeedListStore(db)
	if err != nil {
		t.Fatal("Expected no error")
	}

	list.Add(defaultFeedsList[0])

	feedList, err := list.ListAll()
	if err != nil {
		t.Fatal("Expected no error")
	}

	if !reflect.DeepEqual(feedList, defaultFeedsList) {
		t.Fatal("Expected to still have the original list of feeds when re-adding an existing one")
	}
}

func TestAddingANewEntryAddsItToTheList(t *testing.T) {
	db := newTestDB()
	list, err := NewFeedListStore(db)
	if err != nil {
		t.Fatal("Expected no error")
	}

	list.Add(Feed{
		Title:       "New News",
		Description: "The newest news out there",
		URL:         "http://feed.newynews/news/uk/rss.xml",
	})

	feedList, err := list.ListAll()
	if err != nil {
		t.Fatal("Expected no error")
	}

	found := false
	for _, feed := range feedList {
		if feed.URL == "http://feed.newynews/news/uk/rss.xml" {
			found = true
		}
	}
	if !found {
		t.Fatal("Expected new entry to be in the list")
	}

	if len(feedList) != len(defaultFeedsList)+1 {
		t.Fatal("Expected only one new entry to be added but added ", len(feedList)-len(defaultFeedsList))
	}
}

func TestCanFetchEntryByID(t *testing.T) {
	db := newTestDB()
	list, err := NewFeedListStore(db)
	if err != nil {
		t.Fatal("Expected no error")
	}

	newFeed := Feed{
		Title:       "New News",
		Description: "The newest news out there",
		URL:         "http://feed.newynews/news/uk/rss.xml",
	}
	newID, err := list.Add(newFeed)
	if err != nil {
		t.Fatal("Expected no error")
	}

	storedFeed, err := list.GetByID(newID)
	if err != nil {
		t.Fatal("Expected no error")
	}

	if storedFeed.URL != newFeed.URL {
		t.Fatal("Expected new entry to be in the list")
	}
}

func newTestDB() *bolt.DB {
	tmpfile, err := ioutil.TempFile("", "test.db")
	if err != nil {
		panic(err)
	}
	db, err := bolt.Open(tmpfile.Name(), 0600, nil)
	if err != nil {
		panic(err)
	}

	return db
}
