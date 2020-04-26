// Sample program to show how to read a stack trace when it packs values.
package main

func main() {
	example(true, false, true, 25)
}

//go:noinline
func example(b1, b2, b3 bool, i uint8) {
	panic("Want stack trace")
}

/*
	panic: Want stack trace
	goroutine 1 [running]:
	main.example(0xc019010001)
		stack_trace/example2/example2.go:13 +0x39
	main.main()
		stack_trace/example2/example2.go:8 +0x29
--------------------------------------------------------------------------------
