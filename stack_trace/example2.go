// Sample program to show how to read a stack trace when it packs values.
package main

func main() {
	example(true, false, true, 25)
}

//go:noinline
func example(b1, b2, b3 bool, i uint8) {
	panic("Want stack trace")
}
