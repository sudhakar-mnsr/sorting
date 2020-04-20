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

// Lock prevents a bike from moving
func (bike) Lock() {
   fmt.Println("Locking the bike")
}

// Unlock allows bike to be moved.
func (bike) Unlock() {
   fmt.Println("Unlocking the bike")
}

func main() {
   // Declare variables of the MoveLocker and Mover interfaces set to 
   // their zero value.
   var ml MoveLocker
   var m Mover

   // Create a value of type bike and assign the value to the MoveLocker
   // interface value
   ml = bike{}
   ml = m // allowed

   // It is important to note that the type assertion syntax provides a 
   // way to state what type of value is stored inside the interface. 
   // This is more powerful from a language and readability standpoint,
   // than using a casting syntax, like in other languages.
   b := m.(bike)
   ml = b

   // another form without panic
   if (b1, ok := m.(bike)); ok {
       fmt.Println(ok)
   }
}  
