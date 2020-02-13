package main

import (
	"fmt"
	"io"
	"os"
)

type alphaReader struct {
	src io.Reader
}

func NewAlphaReader(source io.Reader) *alphaReader {
	return &alphaReader{source}
}
