package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var PORT string
var yang = "higgyhiggy"

// takes port number as imput we pick
func main() {
	var port int
	flag.IntVar(&port, "p", 80, "specify port to use.  defaults to 8000.")
	flag.Parse()
	PORT = strconv.Itoa(port)
	fmt.Printf("Host started on port %s\n", PORT)
	//prints out when we hit that port along with the path of it /yoyo/lol
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		answer := req.Header.Get("ying")
		if yang != answer {
			fmt.Println("Access Denied ", req.RemoteAddr)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Acess Denied"))
			return
		}
		w.Write([]byte("Acess GRANTED!!!"))
		fmt.Println(getIP(req), "---> requested this host ")

		ua := req.Header.Get("User-Agent")
		//tu := req.Header.Get("host")

		println(ua)

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
