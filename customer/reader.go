package customer

import (
	"bufio"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

const (
	readingError string = "error: reading impossible"
	emptyError   string = "error: string empty"
)
