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
