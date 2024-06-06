package main

import (
	"fmt"
	"github.com/nade-harlow/go-test-exercise/section3/httpServer/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.HelloHandler)

	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
