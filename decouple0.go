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

// Data is the structure of the data we are copying
type Data struct {
   Line string
}

// Xenia is a system we need to pull data from.
type Xenia struct {
   Host string
   Timeout time.Duration
}

