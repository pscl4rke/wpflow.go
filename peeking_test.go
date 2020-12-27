package main

import "testing"

func bah(src []string) chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, x := range src {
			out <- x
		}
	}()
	return out
}

func assertStringEqual(t *testing.T, value1 string, value2 string) {
	if value1 != value2 {
		t.Errorf("%q != %q", value1, value2)
	}
}

func TestPairs(t *testing.T) {
	input := []string{"foo", "bar", "baz"}
	pairs := PeekStrings(bah(input))
	pair1 := <-pairs
	assertStringEqual(t, pair1.Current, "foo")
	assertStringEqual(t, pair1.Next, "bar")
	pair2 := <-pairs
	assertStringEqual(t, pair2.Current, "bar")
	assertStringEqual(t, pair2.Next, "baz")
	pair3 := <-pairs
	assertStringEqual(t, pair3.Current, "baz")
	assertStringEqual(t, pair3.Next, "")
}
