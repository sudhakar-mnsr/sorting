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
