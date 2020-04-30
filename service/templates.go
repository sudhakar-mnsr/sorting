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
