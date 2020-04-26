package search

import "html/template"

// Options provides the search options for performing searches.
type Options struct {
	Term  string
	CNN   bool
	NYT   bool
	BBC   bool
	First bool
}

// Result represents a search result that was found.
type Result struct {
	Engine  string
	Title   string
	Link    string
	Content string
}

// TitleHTML fixes encoding issues.
func (r *Result) TitleHTML() template.HTML {
	return template.HTML(r.Title)
}

// ContentHTML fixes encoding issues.
func (r *Result) ContentHTML() template.HTML {
	return template.HTML(r.Content)
}

// Searcher declares an interface used to leverage different
// search engines to find results.
type Searcher interface {
	Search(uid string, term string, found chan<- []Result)
}

// Submit uses goroutines and channels to perform a search against the
// feeds concurrently.
func Submit(uid string, options Options) []Result {
	searchers := make(map[string]Searcher)

	// Create a CNN Searcher if checked.
	if options.CNN {
		searchers["cnn"] = NewCNN()
	}

	// Create a NYT Searcher if checked.
	if options.NYT {
		searchers["nyt"] = NewNYT()
	}

	// Create a BBC Searcher if checked.
	if options.BBC {
		searchers["bbc"] = NewBBC()
	}
