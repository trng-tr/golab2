package customer

import (
	"bufio"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

const errorMgs string = "error: reading impossible"
