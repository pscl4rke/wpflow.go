
package main

import (
	"bytes"
	"testing"
)

func TestShortParagraphs(t *testing.T) {
	input := "one\n\ntwo\n\nthree\n"
	expected := input
	source := bytes.NewBufferString(input)
	dest := bytes.Buffer{}
	wpflow(source, &dest)
	if result := dest.String(); result != expected {
		t.Errorf("Invalid %q", result)
	}
}

func TestCodeSnippets(t *testing.T) {
	input := "one\n\n    foo\n    bar\n\nthree\n"
	expected := input
	source := bytes.NewBufferString(input)
	dest := bytes.Buffer{}
	wpflow(source, &dest)
	if result := dest.String(); result != expected {
		t.Errorf("Invalid %q", result)
	}
}
