// Sample program to show how stacks grow/change.
package main

// Number of elements to grow each stack frame.
// Run with 1 and then with 1024
const size = 1

// main is the entry point for the application.
func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}
