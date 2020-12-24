
// wpflow reformats markdown for better wordpress rendering
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	wpflow(os.Stdin, os.Stdout)
}

func wpflow(src io.Reader, dest io.Writer) {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintln(dest, line)
	}
}
