package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, "+os.Getenv("ENV"))
	})
	http.ListenAndServe(":8080", nil)
}
