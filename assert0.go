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
