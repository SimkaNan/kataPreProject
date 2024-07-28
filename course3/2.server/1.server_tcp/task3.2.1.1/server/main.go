// Код сервера
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

var (
	// канал для всех входящих клиентов
	entering = make(chan client)
	// канал для сообщения о выходе клиента
	leaving = make(chan client)
	// канал для всех сообщений
	messages = make(chan string)
	//канал через который отправляется ник автора
	nameSender = make(chan string)
)

func main() {
	listener := setListn()

	go broadcaster()
	acceptConn(listener)
}

func setListn() net.Listener {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return listener
}

func acceptConn(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

// broadcaster рассылает входящие сообщения всем клиентам
// следит за подключением и отключением клиентов
func broadcaster() {
	// здесь хранятся все подключенные клиенты
	clients := make(map[client]bool)

	go clEnter(clients)
	go clLeft(clients)

	var who string

	for {
		msg := <-messages

		select {
		case tmp := <-nameSender:
			who = tmp
		default:
			who = "..."
		}

		for client, check := range clients {
			if check == true && who != client.name {
				_, err := client.conn.Write([]byte(msg))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func clEnter(clients map[client]bool) {
	for {
		clients[<-entering] = true
	}
}

func clLeft(clients map[client]bool) {
	for {
		clients[<-leaving] = false
	}
}

// handleConn обрабатывает входящие сообщения от клиента
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{conn, who, ch}

	ch <- "You are " + who
	messages <- who + " has arrived"
	nameSender <- who
	entering <- cli

	readMsg(conn, who)

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

// функция для принятия и отправки сообщений
func readMsg(conn net.Conn, who string) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println(err)
			continue
		}
		msg := string(buffer[:n])
		if msg == "/quit" {
			break
		}

		messages <- who + ": " + msg
		nameSender <- who
	}
}

// clientWriter отправляет сообщения текущему клиенту
func clientWriter(conn net.Conn, ch <-chan string) {
	for {
		_, err := conn.Write([]byte(<-ch))
		if err != nil {
			fmt.Println(err)
		}
	}
}
