// Sample program to show how to read a stack trace.
package main

func main() {
	example(make([]string, 2, 4), "hello", 10)
}


//go:noinline
func example(slice []string, str string, i int) {
	panic("Want stack trace")
}

/*
	panic: Want stack trace
	goroutine 1 [running]:
	main.example(0xc000042748, 0x2, 0x4, 0x106abae, 0x5, 0xa)
		stack_trace/example1/example1.go:13 +0x39
	main.main()
		stack_trace/example1/example1.go:8 +0x72
	--------------------------------------------------------------------------------
