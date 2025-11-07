package main

import (
	"fmt"
	"http-from-tcp/internal/request"
	"log"
	"net"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		// fmt.Println("Accepted connection from", conn.RemoteAddr())

		requestLine, err := request.RequestFromReader(conn)

		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		fmt.Println("Request line:")
		fmt.Println("- Method:", requestLine.RequestLine.Method)
		fmt.Println("- Target:", requestLine.RequestLine.RequestTarget)
		fmt.Println("- Version:", requestLine.RequestLine.HttpVersion)
		// fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
	}
}
