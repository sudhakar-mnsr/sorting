package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rows := []string{
		"The quick brown fox",
		"jumps over the lazy dog",
	}

	fout, err := os.Create("./filewrite.data")
	writer := bufio.NewWriter(fout)
