package search

import "log"

var nytFeeds = []string{
	"http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
	"http://rss.nytimes.com/services/xml/rss/nyt/US.xml",
	"http://rss.nytimes.com/services/xml/rss/nyt/Politics.xml",
	"http://rss.nytimes.com/services/xml/rss/nyt/Business.xml",
}

// NYT provides support for NYT searches.
type NYT struct{}
