package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var PORT string
var Key = "higgyhiggy"

// takes port number as imput we pick
func main() {
	var port int
	flag.IntVar(&port, "p", 80, "specify port to use.  defaults to 8000.")
	flag.Parse()
	PORT = strconv.Itoa(port)

	//prints out when we hit that port along with the path of it /yoyo/lol
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		remoteKey := req.Header.Get("X-Secret-Key")
		if Key != remoteKey {
			fmt.Println("Access Denied ", req.RemoteAddr)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Acess Denied"))
			return
		}
		fmt.Println("New Connection ", req.RemoteAddr)

		ua := req.Header.Get("User-Agent")
		//tu := req.Header.Get("host")
		println(getIP(req))
		println(ua)
		fmt.Printf("Server started on port %s\n", PORT)
	})

	err := http.ListenAndServe(":"+PORT, nil)
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
