package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)
var PORT string

// takes port number as imput we pick
func main() {
	var port int
	flag.IntVar(&port, "p", 80, "specify port to use.  defaults to 8000.")
	flag.Parse()
	PORT = strconv.Itoa(port)


	//prints out when we hit that port along with the path of it /yoyo/lol
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello wolrld")
		ua := req.Header.Get("User-Agent")
		//tu := req.Header.Get("host")
		println(getIP(req))
		println(ua)
		//println("--->", os.Args[1], req.URL.String(), "\n")
		fmt.Printf("Server started on port %s\n", PORT)
	})
	// port that was passed via user that we will listen and server
	//http.ListenAndServeTLS(":"+os.Args[1], "https-server.crt", "https-server.key", nil)
	err := http.ListenAndServe(":"+os.Args[1], nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// to get ip address
func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
