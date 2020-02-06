package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

var PORT = "8000"
var index = 1
var yang = "higgyhiggy"

func newMultipleHostReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.Header.Add("ying", yang)
		//target is the host we will send the request to
		// round robin load balancer
		//target := targets[rand.Int()%len(targets)]
		target := targets[index%2]
		index++
		/*type URL struct {
		    Scheme     string
		    Opaque     string    // encoded opaque data
		    User       *Userinfo // username and password information
		    Host       string    // host or host:port
		    Path       string    // path (relative paths may omit leading slash)
		    RawPath    string    // encoded path hint (see EscapedPath method); added in Go 1.5
		    ForceQuery bool      // append a query ('?') even if RawQuery is empty; added in Go 1.7
		    RawQuery   string    // encoded query values, without '?'
		    Fragment   string    // fragment for references, without '#'
		}*/
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
		println(getIP(req) + "---> accessed rproxy server!!!")

	}
	/* // Director must be a function which modifies
	   // the request into a new request to be sent
	   // using Transport. Its response is then copied
	   // back to the original client unmodified.
	   // Director must not access the provided Request
	   // after returning.
	   Director func(*http.Request)

	   // The transport used to perform proxy requests.
	   // If nil, http.DefaultTransport is used.
	   Transport http.RoundTripper
	*/
	return &httputil.ReverseProxy{
		Director: director,
		Transport: &http.Transport{
			/* Transport is an implementation of RoundTripper that supports HTTP, HTTPS, and HTTP proxies (for either HTTP or HTTPS with CONNECT).

			By default, Transport caches connections for future re-use. This may leave many open connections when accessing many hosts.
			This behavior can be managed using Transport's CloseIdleConnections method and the MaxIdleConnsPerHost and DisableKeepAlives fields. */

			Proxy: func(req *http.Request) (*url.URL, error) {
				/*   // Proxy specifies a function to return a proxy for a given
				// Request. If the function returns a non-nil error, the
				// request is aborted with the provided error.
				//
				// The proxy type is determined by the URL scheme. "http",
				// "https", and "socks5" are supported. If the scheme is empty,
				// "http" is assumed.
				//
				// If Proxy is nil or returns a nil *URL, no proxy is used.*/
				return http.ProxyFromEnvironment(req)
			},
			Dial: func(network, addr string) (net.Conn, error) {
				/*// Dial specifies the dial function for creating unencrypted TCP connections.
				  //
				  // Dial runs concurrently with calls to RoundTrip.
				  // A RoundTrip call that initiates a dial may end up using
				  // a connection dialed previously when the earlier connection
				  // becomes idle before the later Dial completes.
				  //
				  // Deprecated: Use DialContext instead, which allows the transport
				  // to cancel dials as soon as they are no longer needed.
				  // If both are set, DialContext takes priority.*/
				conn, err := (&net.Dialer{
					/* A Dialer contains options for connecting to an address.

					The zero value for each field is equivalent to dialing without that option.
					 Dialing with the zero value of Dialer is therefore equivalent to just calling the Dial function. */
					// Timeout is the maximum amount of time a dial will wait for
					// a connect to complete. If Deadline is also set, it may fail
					// earlier.
					//
					// KeepAlive specifies the interval between keep-alive
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial(network, addr)
				if err != nil {
					println("Error during DIAL:", err.Error())
				}
				return conn, err
			}, // TLSHandshakeTimeout specifies the maximum amount of time waiting to
			// wait for a TLS handshake. Zero means no timeout.
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}

func main() {
	fmt.Println(os.Getenv("rproxy"))
	proxy := newMultipleHostReverseProxy([]*url.URL{
		{
			Scheme: "http",
			Host:   "localhost:9091",
		},
		{
			Scheme: "http",
			Host:   "localhost:9092",
		},
	})
	log.Fatal(http.ListenAndServeTLS(":"+PORT, "server.crt", "server.key", proxy))
}
func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
