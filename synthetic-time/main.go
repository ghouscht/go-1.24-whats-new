package main

import (
	"errors"
	"fmt"
	"time"
)

// Read reads a value from the input channel and returns it. Read times out after
// 10s.
func Read(in chan int) (int, error) {
	select {
	case v := <-in:
		return v, nil
	case <-time.After(10 * time.Second):
		return 0, errors.New("timeout")
	}
}

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	val, err := Read(ch)
	fmt.Printf("read val=%v, err=%v\n", val, err)
}
