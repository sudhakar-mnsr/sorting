// Calculate sum of all multiple of 3 and 5 less than MAX value.
// See https://projecteuler.net/problem=1
package main

import (
	"fmt"
	"sync"
)

const MAX = 1000
const workers = 2

func main() {
	values := make(chan int)
	result := make(chan int, workers)
	var wg sync.WaitGroup
