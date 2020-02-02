package main

import (
	"fmt"
	"net"
)

var ConnSignal chan string = make(chan string)
var connections []net.Conn

func main() {
	ln, _ := net.Listen("tcp", ":8080")

	for {
		go Session(ln)
		fmt.Println(<-ConnSignal)
	}

}

func Session(ln net.Listener) {
	conn, _ := ln.Accept() //Waits
	defer conn.Close()
	connections = append(connections, conn)
	ConnSignal <- "Connection Established"

	for {
		buf := make([]byte, 1024)
		conn.Read(buf)
		for _, c := range connections {
			c.Write(buf)
			fmt.Fprintf(c, "here I be")
		}
	}

}