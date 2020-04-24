// Sample program to show how to read a stack trace
package main

import (
   "fmt"
   "runtime/debug"
)

func main() {
   example(make([]string, 2,4), "hello", 10)
}

func example(slice[]string, str string, i int) {
//   panic("Want stack trace")
   fmt.Println("Stacktrace from Panic\n" + string(debug.Stack()))
}

