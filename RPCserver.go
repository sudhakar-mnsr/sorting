package main 
 
import ( 
   "fmt" 
   "net" 
   "net/rpc" 
   "os" 
   "sharedRPC" 
)

type MyInterface int 

func (t *MyInterface) Add(arguments *sharedRPC.MyInts, reply *int) error {
   s1 := 1
   s2 := 1
   if arguments.S1 == true {
      s1 = -1
   }
   if arguments.S2 == true {
      s2 = -1
   }

   *reply = s1 * int(arguments.S1) + s2 * int(arguments.S2)
   return nil
}
