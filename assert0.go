package main

import "fmt"

// Mover provides support for moving things.
type Mover interface {
   Move()
}

// Locker provides support for locking and unlocking things
type Locker interface {
   Lock()
   Unlock()
}

// MoveLocker provides support for moving and locking things.
type MoveLocker interface {
   Mover
   Locker
}

// bike represents a concrete type for the example
type bike struct{}

// Move can change the position of a bike
func (bike) Move() {
   fmt.Println("Moving the bike")
}
