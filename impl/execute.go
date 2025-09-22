package impl

import (
	"fmt"
	"net/http"
)

func Execute() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
