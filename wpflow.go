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
	for pair := range PeekStrings(linesIn(src)) {
		fmt.Fprint(dest, pair.Current)
		fmt.Fprint(dest, separatorBetween(pair))
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

func separatorBetween(pair StringPair) string {
	if pair.Current == "" {
		return "\n"
	}
	if pair.Next == "" {
		return "\n"
	}
	if pair.Next[0] == ' ' {
		return "\n"
	}
	return " "
}
