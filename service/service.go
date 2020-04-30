// Package service maintains the logic for the web service.
package service

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/braintree/manners"
)

// init binds the routes and handlers for the web service.
func init() {

	// Setup a route for our static files.
	//
	// Because our static directory is set as the root of the FileSystem,
	// we need to strip off the /static/ prefix from the request path
	// before searching the FileSystem for the given file.
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Setup a route for the home page.
	http.HandleFunc("/search", handler)
}
