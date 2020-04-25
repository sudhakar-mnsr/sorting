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

