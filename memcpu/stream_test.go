// go test -bench . -benchmem -memprofile p.out -gcflags "-newescape=false -m=2"

// Tests to see how each algorithm compare.
package main

import (
	"bytes"
	"testing"
)
