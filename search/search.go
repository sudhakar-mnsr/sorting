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
