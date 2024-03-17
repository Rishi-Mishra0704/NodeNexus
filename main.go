package main

import (
	"log"
	"net/http"
)

func main() {
	// Start your server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
