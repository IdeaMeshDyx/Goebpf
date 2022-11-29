package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Path
func indexHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(rw http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(rw, "Header[%q] = %q\n", k, v)
	}
}
