package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func handleConnection(conn net.Conn, file *os.File) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		request, err := http.ReadRequest(reader)
		if err != nil {
			fmt.Println("Error reading request:", err)
			break
		}

		if request.Method != http.MethodGet {
			response := http.Response{
				StatusCode: 404,
				Body:       io.NopCloser(strings.NewReader("not found")),
			}

			err = response.Write(conn)
			if err != nil {
				fmt.Println("Error writing response:", err)
				break
			}
		}

		response := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Body:       file,
			Header:     make(http.Header),
		}
		response.Header.Set("Content-Type", "text/plain")

		err = response.Write(conn)
		if err != nil {
			fmt.Println("Error writing response:", err)
			break
		}
	}
}

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()
// 	buf := make([]byte, 1024)
// 	for {
// 		n, err := conn.Read(buf)
// 		if err == io.EOF {
// 			return
// 		} else if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		fmt.Println(string(buf[:n]))
// 	}
// }

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	file, err := os.Open("get.html")
	if err != nil {
		fmt.Println("Невозможно прочитать файл:", err)
	}

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn, file)
	}
}
