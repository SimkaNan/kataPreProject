package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	request, err := http.ReadRequest(reader)
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	if request.Method == http.MethodGet && request.URL.Path == "/" {
		file := openHtml()

		fmt.Println(request.URL.Path)

		body, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		file.Close()

		cntLngth := strconv.Itoa(len(string(body)))
		status := "HTTP/1.1 200 OK\nContent-Length: " + cntLngth + "\n"
		resp := []byte(status)
		resp = append(resp, body...)

		_, err = conn.Write(resp)
		if err != nil {
			fmt.Println(err)
		}

		log.Println("GET-запрос успешно обработан")
	} else {
		response := http.Response{
			StatusCode: 404,
			ProtoMajor: 1,
			ProtoMinor: 1,
		}

		err = response.Write(conn)
		if err != nil {
			fmt.Println("Error writing response:", err)
			return
		}

	}
}

func openHtml() *os.File {
	file, err := os.Open("get.html")
	if err != nil {
		fmt.Println("Невозможно прочитать файл:", err)
	}

	return file
}

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}
