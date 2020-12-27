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
	for line := range linesIn(src) {
		fmt.Fprintln(dest, line)
	}
}

func linesIn(src io.Reader) chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		scanner := bufio.NewScanner(src)
		for scanner.Scan() {
			out <- scanner.Text()
		}
	}()
	return out
}
