package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// takes port number as imput we pick
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <port>", os.Args[0])
	}
	if _, err := strconv.Atoi(os.Args[1]); err != nil {
		log.Fatalf("Invalid port: %s (%s)\n", os.Args[1], err)
	}
	//prints out when we hit that port along with the path of it /yoyo/lol
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello wolrld")
		println("--->", os.Args[1], req.URL.String())
	})
	// port that was passed via user that we will listen and server
	http.ListenAndServe(":"+os.Args[1], nil)
}
