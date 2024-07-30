package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//
//func main() {
//	http.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "<h1>Hello from Docker container!</h1>")
//	})
//
//	err := http.ListenAndServe("localhost:8080", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
