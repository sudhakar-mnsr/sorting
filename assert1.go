// Sample program demonstrating that type assertions are a runtime and
// not compile time construct.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// car represents something you drive
type car struct{}

// String implements the fmt.Stringer interface
func (car) String() string {
   return "Vroom!"
}
