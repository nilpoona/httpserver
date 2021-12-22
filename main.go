package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()
			reader := newRequestReader(c)
			requestLine, err := reader.ReadLine()
			if err != nil {
				log.Fatal(err)
				return
			}
			method, uri, version := parseRequestLine(string(requestLine))
			fmt.Printf("method=%s, uri=%s, version=%s\n", method, uri, version)
			headers, err := reader.ReadHeaders()
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Printf("headers=%+v\n", headers)
			message, err := reader.ReadMessageBody(headers)
			if err != nil {
				log.Fatal(err)
				return
			}
			// TODO parse request message
			fmt.Printf("message=%+v\n", message)
		}(conn)
	}
}
