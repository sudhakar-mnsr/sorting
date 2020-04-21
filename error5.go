// Sample program to show how wrapping errors work.
package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// AppError represents a custom error type.
type AppError struct {
	State int
}
