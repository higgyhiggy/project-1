package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	var conn net.Conn
	var err error
	for {
		conn, err = net.Dial("tcp", "localhost:8080")
		if err == nil {
			break
		}
	}

	go Lisenter(conn)
	Writer(conn)
}

func Lisenter(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		conn.Read(buf)
		if string(buf) == "Hello" {

		}
		fmt.Print(string(buf))
	}
}

func Writer(conn net.Conn) {
	for {
		r := bufio.NewReader(os.Stdin)
		text, _ := r.ReadString('\n')

		conn.Write([]byte(text))
	}
}