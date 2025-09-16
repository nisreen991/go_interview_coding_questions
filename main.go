package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	exampleTimeout()

}

func exampleTimeout() {
	ctx := context.Background()

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("Calling the API")

	case <-ctxWithTimeout.Done():
		fmt.Println("Time out exceeded", ctxWithTimeout.Err())
	}
}
