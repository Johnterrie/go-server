package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)

	fmt.Println("Starting the Service on :8080")

	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println("Server failed", err)
	}
}
