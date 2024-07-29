// Код клиента
package main

import (
	"fmt"
	"net"
)

func main() {
	// Устанавливаем соединение с сервером
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		return
	}
	defer conn.Close()

	// // Отправка данных на сервер
	successReq(conn)

	conn, err = net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		return
	}
	defer conn.Close()

	failedReq(conn)
}

func successReq(conn net.Conn) {
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		return
	}

	fmt.Printf("%s\n\n", string(buf[:n]))
}

func failedReq(conn net.Conn) {
	buf := make([]byte, 1024)

	fmt.Fprintf(conn, "GET /check.org/like HTTP/1.1\r\n\r\n")

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу:", err)
		return
	}

	fmt.Println(string(buf[:n]))
}
