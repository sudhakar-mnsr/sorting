// Sample test to show how to write a basic unit test.
package search

import "testing"

var final []Result

// BenchmarkRssSearch provides support for profiling the search.
func BenchmarkRssSearch(b *testing.B) {
	var result []Result
	var err error

	for i := 0; i < b.N; i++ {
		result, err = rssSearch("1", "trump", "nyt", "http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml")
		if err != nil {
			b.FailNow()
		}
	}

	final = result
}
