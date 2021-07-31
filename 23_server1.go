package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
* A basic server
 */
func main() {

	http.HandleFunc("/", handler) //req handler
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from IP :-> ", r.RemoteAddr)
	fmt.Fprintf(w, "Request path->: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Request method->: %s\n", r.Method)
	fmt.Fprintf(w, "Request headers->:\n")
	for k, v := range r.Header{
         fmt.Fprintf(w, "\t %s: %s\n", k, v)
	}
}
