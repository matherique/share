package main

import (
	"fmt"
	"log/slog"
	"net"
	"strings"
)

const (
	okResponse = "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\n"
)

func main() {
	net, err := net.Listen("tcp", ":8080")

	if err != nil {
		slog.Error("fail to create listener", "error", err)
		return
	}
	defer net.Close()

	msg := make(chan string)
	go server(msg)

	for {
		conn, err := net.Accept()
		if err != nil {
			slog.Error("fail to create listener", "error", err)
			continue
		}

		go client(conn, msg)

	}
}

func server(msg chan string) {
	for {
		data := strings.SplitN(<-msg, "\r\n\r\n", 2)

		if len(data) == 1 {
			continue
		}

		body := data[1]

		fmt.Println(body)
	}
}

func client(conn net.Conn, msg chan<- string) {
	defer conn.Close()
	buff := make([]byte, 1024)

	n, err := conn.Read(buff)
	if err != nil {
		return
	}

	msg <- string(buff[0:n])
	conn.Write([]byte(okResponse))
}
