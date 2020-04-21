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

// Error implements the error interface.
func (c *AppError) Error() string {
	return fmt.Sprintf("App Error, State: %d", c.State)
}

func main() {

	// Make the function call and validate the error.
	if err := firstCall(10); err != nil {

		// Use type as context to determine cause.
		switch v := errors.Cause(err).(type) {
		case *AppError:

			// We got our custom error type.
			fmt.Println("Custom App Error:", v.State)
