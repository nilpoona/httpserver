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
			reader := newRequestReader(c)
			requestLine, err := reader.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			method, uri, version := parseRequestLine(string(requestLine))
			fmt.Printf("method=%s, uri=%s, version=%s", method, uri, version)
			c.Close()
		}(conn)
	}
}