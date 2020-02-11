package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	curr "currency/lib"
)

var currencies = curr.Load("../data.csv")



