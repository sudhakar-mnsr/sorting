// Sample program that takes a stream of bytes and looks for the bytes
// “elvis” and when they are found, replace them with “Elvis”. The code
// cannot assume that there are any line feeds or other delimiters in the
// stream and the code must assume that the stream is of any arbitrary length.
// The solution cannot meaningfully buffer to the end of the stream and
// then process the replacement.
package main

import (
	"bytes"
	"fmt"
	"io"
)

// data represents a table of input and expected output.
var data = []struct {
	input  []byte
	output []byte
}{
	{[]byte("abc"), []byte("abc")},
	{[]byte("elvis"), []byte("Elvis")},
	{[]byte("aElvis"), []byte("aElvis")},
	{[]byte("abcelvis"), []byte("abcElvis")},
	{[]byte("eelvis"), []byte("eElvis")},
	{[]byte("aelvis"), []byte("aElvis")},
	{[]byte("aabeeeelvis"), []byte("aabeeeElvis")},
	{[]byte("e l v i s"), []byte("e l v i s")},
	{[]byte("aa bb e l v i saa"), []byte("aa bb e l v i saa")},
	{[]byte(" elvi s"), []byte(" elvi s")},
	{[]byte("elvielvis"), []byte("elviElvis")},
	{[]byte("elvielvielviselvi1"), []byte("elvielviElviselvi1")},
	{[]byte("elvielviselvis"), []byte("elviElvisElvis")},
}

// assembleInputStream combines all the input into a
// single stream for processing.
func assembleInputStream() []byte {
	var in []byte
	for _, d := range data {
		in = append(in, d.input...)
	}
	return in
}

// assembleOutputStream combines all the output into a
// single stream for comparing.
func assembleOutputStream() []byte {
	var out []byte
	for _, d := range data {
		out = append(out, d.output...)
	}
	return out
}

func main() {
	var output bytes.Buffer
	in := assembleInputStream()
	out := assembleOutputStream()

	find := []byte("elvis")
	repl := []byte("Elvis")

	fmt.Println("=======================================\nRunning Algorithm One")
	output.Reset()
	algOne(in, find, repl, &output)
	matched := bytes.Compare(out, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, in, out, output.Bytes())


	fmt.Println("=======================================\nRunning Algorithm Two")
	output.Reset()
	algTwo(in, find, repl, &output)
	matched = bytes.Compare(out, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, in, out, output.Bytes())
}

// algOne is one way to solve the problem.
func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {

	// Use a bytes Buffer to provide a stream to process.
	input := bytes.NewBuffer(data)

	// The number of bytes we are looking for.
	size := len(find)

	// Declare the buffers we need to process the stream.
	buf := make([]byte, size)
	end := size - 1

	// Read in an initial number of bytes we need to get started.
	if n, err := io.ReadFull(input, buf[:end]); err != nil {
		output.Write(buf[:n])
		return
	}
