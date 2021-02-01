# Testing

Run 
```
go test ./...
```
from the root of the project

# Running

You can either:
- Build it and run it by running this from the root:
```
go run main.go
```
- Run the prebuilt binary that I've included for demo purposes by running this from root:
```
./feed-reader
```

Either of these will run a server on port `:8080`

# Storage

Because this is a demo app I opted for BoltDB as an in-memory file backed store. If you want to reset the storage to the initial state then delete the feed-reader.db file.

# Example usage

The endpoints have more documentation for the expected request and response formats, but this is a quick list of things you can try on the API:

## Listing available feeds
```
curl -v -H "Content-Type: application/json" "localhost:8080/api/feeds"
```
## Reading an individual feed
```
curl -v -H "Content-Type: application/json" "localhost:8080/api/feeds/b1031651-411c-40bb-b269-d247794dfd59"
```
## Reading an individual feed with pagination
```
curl -v -H "Content-Type: application/json" "localhost:8080/api/feeds/b1031651-411c-40bb-b269-d247794dfd59?per_page=5&offset=0"
```
## Reading an individual feed with category filtering
```
curl -s -H "Content-Type: application/json" localhost:8080/api/feeds/3f58c0ec-88e5-49e0-b147-75f906ba135e?category=Culture
```
## Adding a new feed
```
curl -v -X POST -H "Content-Type: application/json" --data '{"url": "https://www.theguardian.com/uk/rss"}' "localhost:8080/api/feeds"
```

If you'd like to see the documentation outside the code you can run `godoc` in the root directory and navigate your browser to `http://localhost:6060/pkg/feed-reader/app/`, you can also check out the other packages there

# Design Choices
- I chose not to use a testing framework because I wanted to keep the code as simple as possible and to not lean towards any particular testing style
- I chose to use as few libraries as possible in order to show more implementation, in real life I would have likely used a few more, particularly for the RSS parsing but I thought the code would be less interesting if I just used something off the shelf.
- I chose to store the list of registered feeds in a db saved to a local file. This is not ok for cloud based service because it should be statless. The app would work just as fine with something like PG but I wanted to avoid any setup complications and have it all easily runnable as possible.
- I chose to have a hardcoded list of the example feeds in the requirements. I could have prepopulated the DB file or had a step on running where I would make the user populate them. Either is fine, but again I wanted to make this super easy to check from the get go.
- Some feeds seem somewhat custom. The exercise mentions showing an image for each article. Some feeds do provide images in the feed but the RSS standard doesn't seem to mention them so I chose to leave this part out for now. It wouldn't be too hard to provide a customised parser based on hostnames but it's not a one-size fits all solution. I may have misunderstood the format, if so it would be straightforward to add them in.
- I approached the code from an inside-out TDD strategy. I tend to favor this when I have no structure to the codebase yet and I also wanted to focus my testing on the core logic of the app. I started running out of time at the end but would have liked to add an integration test per endpoint just to make sure everything is wired up properly.
- Everything under a `Note:` comment is meant for the exercise review and I would not normally add this kind of comment in a real codebase.
