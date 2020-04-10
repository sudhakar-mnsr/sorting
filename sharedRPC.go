package sharedRPC

type MyInts struct {
   A1, A2 uint
   S1, S2 bool
}

type MyInterface interface {
   Add(arguments *MyInts, reply *int) error
   Sub(arguments *MyInts, reply *int) error
}

