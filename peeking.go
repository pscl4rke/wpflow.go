package main

type StringPair struct {
	Current string
	Next string
}

func PeekStrings(src chan string) chan StringPair {
	out := make(chan StringPair)
	go func() {
		defer close(out)
		current, okay := <-src
		if !okay {
			return
		}
		for next := range src {
			out <- StringPair{current, next}
			current = next
		}
		out <- StringPair{current, ""}
	}()
	return out
}
