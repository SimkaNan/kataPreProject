package main

import (
	"log"
	"net/http"
)

type MyHandler struct{}

func main() {
	var h MyHandler

	err := http.ListenAndServe("localhost:8080", h)
	if err != nil {
		log.Fatal(err)
	}
}

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hello world")
	_, err := res.Write(data)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Запрос успешно обработан")
}
