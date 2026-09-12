package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "Hello, world!")
	})

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
