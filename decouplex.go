// Sample program demonstrating being more precise with API design.
package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// =============================================================================

// Data is the structure of the data we are copying.
type Data struct {
	Line string
}

// =============================================================================

// Puller declares behavior for pulling data.
type Puller interface {
	Pull(d *Data) error
}

// Storer declares behavior for storing data.
type Storer interface {
	Store(d *Data) error
}

// =============================================================================

// Xenia is a system we need to pull data from.
type Xenia struct {
	Host    string
	Timeout time.Duration
}
