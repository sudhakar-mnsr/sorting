// Package service : temnplates provides support for using HTML
// based templates for responses.
package service

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

// views contains a map of templates for rendering views.
var views = make(map[string]*template.Template)

// init loads the existing templates for use by routing code.
func init() {
	// In order for the endpoint tests to run this needs to be
	// physically located. Trying to avoid configuration for now.
	pwd, _ := os.Getwd()
	loadTemplate("layout", pwd+"/views/basic-layout.html")
	loadTemplate("search", pwd+"/views/search.html")
	loadTemplate("results", pwd+"/views/results.html")
}
