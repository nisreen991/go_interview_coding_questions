package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	exampleTimeout()
	exampleValue()
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

// example 1
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

func exampleValue() {
	ctx := context.Background()

	ctxWithValue := context.WithValue(ctx, "user", 123)

	if userId, exists := ctxWithValue.Value("user").(int); exists {
		fmt.Println("UserID", userId)
	} else {
		fmt.Println("No user with userid found")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("API Response")

	case <-ctx.Done():
		fmt.Println("Context expired")
		http.Error(w, "Timeout occured", http.StatusRequestTimeout)
		return
	}
}
