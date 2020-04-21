// Sample program to show how the default error type is implemented.
package main

import "fmt"

// http://golang.org/pkg/builtin/#error
type error interface {
	Error() string
}
