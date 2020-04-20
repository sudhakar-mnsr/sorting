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

// cloud represents somewhere you store information
type cloud struct{}

// string implements the fmt.Stringer interface.
func (cloud) String() string {
   return "Big Data!"
}

func main() {
   // Seed the number random generator
   rand.Seed(time.Now().UnixNano())

   // Create a slice of the Stringer interface values.
   mvs := []fmt.Stringer{
            car{},
            cloud{},
   }

