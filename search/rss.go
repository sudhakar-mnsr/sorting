package search

import (
	"encoding/xml"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	gc "github.com/patrickmn/go-cache"
)

// Maintain a cache of retrieved documents. The cache will maintain items for
// fifteen seconds and then marked as expired. This is a very small cache so the
// gc time will be every hour.

const (
	expiration = time.Minute * 15
	cleanup    = time.Hour
)

var cache = gc.New(expiration, cleanup)


var fetch = struct {
	sync.Mutex
	m map[string]*sync.Mutex
}{
	m: make(map[string]*sync.Mutex),
}

type (

	// Item defines the fields associated with the item tag in the RSS document.
	Item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
	}
