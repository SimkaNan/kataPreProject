package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from Docker container!</h1>")
	})

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
