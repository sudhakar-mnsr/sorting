package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type alphaReader struct {
	src io.Reader
}
