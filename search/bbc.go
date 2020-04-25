package search

import "log"

var bbcFeeds = []string{
	"http://feeds.bbci.co.uk/news/rss.xml",
	"http://feeds.bbci.co.uk/news/world/rss.xml",
	"http://feeds.bbci.co.uk/news/politics/rss.xml",
	"http://feeds.bbci.co.uk/news/world/us_and_canada/rss.xml",
}

// BBC provides support for BBC searches.
type BBC struct{}

// NewBBC returns a BBC Searcher value.
func NewBBC() Searcher {
	return BBC{}
}
