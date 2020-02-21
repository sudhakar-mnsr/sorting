// Calculate sum of all multiple of 3 and 5 less than MAX value.
// See https://projecteuler.net/problem=1
package main

import (
	"fmt"
	"sync"
)

const MAX = 1000

func main() {
	values := make(chan int, MAX)
	result := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(2)
