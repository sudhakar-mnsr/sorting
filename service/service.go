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
