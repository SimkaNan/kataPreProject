// Код клиента
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn := connectServer()

	// запустить горутину, которая будет читать все сообщения от сервера и выводить их в консоль
	go clientReader(conn)

	//запустить горутину которая будет читать из stdin и отправлять на сервер
	sendMessage(conn)
}

func connectServer() net.Conn {
	// подключиться к серверу
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return conn
}

// clientReader выводит на экран все сообщения от сервера
func clientReader(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Server shut down")
				return
			}
			fmt.Println(err)
			continue
		}

		fmt.Printf("%s\n", string(buffer[:n]))
	}
}

func sendMessage(conn net.Conn) {
	// читать сообщения от stdin и отправлять их на сервер
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		_, err := conn.Write([]byte(msg))
		if err != nil {
			if msg == "/quit" {
				return
			}

			fmt.Println(err)
			continue
		}
		if msg == "/quit" {
			return
		}
	}

	fmt.Println("хуйня закончена")
}
