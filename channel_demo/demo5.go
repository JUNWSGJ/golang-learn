package main

import (
	"time"
)

type result struct {
	record string
	err error
}

func search(term string) (string, error) {
	time.Sleep(200 *time.Millisecond)
	return "some value", nil
}

func processSearch(term string) error {

	ch := make(chan result)
	go func() {
		record, err := search(term)
		ch <- result{record, err}
	}()

	return nil
}

func main() {

}
