package read

import (
	"bufio"
	"os"
)

/*nous allons utiliser le stream reader
bufio.NewReader qui lit un flux de stream.
bufio.NewReader est robuste que bufio.NewSacanner
*/

var StreamReader = bufio.NewReader(os.Stdin)
