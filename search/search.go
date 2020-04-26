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
