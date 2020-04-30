package service

import (
	"expvar"
	"fmt"
	"html/template"
	"net/http"

	"github.com/ardanlabs/gotraining/topics/go/profiling/project/search"
	"github.com/pborman/uuid"
)

// req keeps track of the number of requests.
var req = expvar.NewInt("requests")

// handler handles the search route processing.
func handler(w http.ResponseWriter, r *http.Request) {
	uid := uuid.New()

	// Add a new counter for monitoring.
	req.Add(1)

	// Capture all the form values.
	fv, options := formValues(r)

	// If this is a post, perform a search.
	var results []search.Result
	if r.Method == "POST" && options.Term != "" {
		results = search.Submit(uid, options)
	}

	// Render the search page.
	markup := render(fv, results)

	// Write the final markup as the response.
	fmt.Fprint(w, string(markup))
}

// formValues extracts the form data.
func formValues(r *http.Request) (map[string]interface{}, search.Options) {
	fv := make(map[string]interface{})
	var options search.Options

	fv["term"] = r.FormValue("term")
	options.Term = r.FormValue("term")
