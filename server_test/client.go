package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}

	// to server
	conn.Write([]byte("Hello Golang"))
}
