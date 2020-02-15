package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var books bytes.Buffer
	books.WriteString("The Great Gatsby\n")
	books.WriteString("1984\n")
	books.WriteString("A Take of Two Cities\n")
	books.WriteString("Les Miserables\n")
	books.WriteString("The Call of the Wild\n")
