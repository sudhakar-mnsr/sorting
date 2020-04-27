// Sample program to show how to use the http trace with a unique Client
// and Transport.
package main

import (
	"log"
	"net/http"
	"net/http/httptrace"
)

// transport is an http.RoundTripper that keeps track of the in-flight
// request and implements hooks to report HTTP tracing events.
type transport struct {
	current *http.Request
}
