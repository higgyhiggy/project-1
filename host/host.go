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
	fmt.Println(os.Getenv("host"))
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <port>", os.Args[0])
	}
	if _, err := strconv.Atoi(os.Args[1]); err != nil {
		log.Fatalf("Invalid port: %s (%s)\n", os.Args[1], err)
	}
	//prints out when we hit that port along with the path of it /yoyo/lol
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello wolrld")
		ua := req.Header.Get("User-Agent")
		//tu := req.Header.Get("host")
		println(getIP(req))
		println(ua)
		println("--->", os.Args[1], req.URL.String(), "\n")
	})
	// port that was passed via user that we will listen and server
	//http.ListenAndServeTLS(":"+os.Args[1], "https-server.crt", "https-server.key", nil)
	err := http.ListenAndServe(":"+os.Args[1], nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
