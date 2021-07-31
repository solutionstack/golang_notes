package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

/**
* A basic server, that keeps track of request count and implements graceful shutdown
 */

var mu sync.Mutex
var count int

func main() {

	server := &http.Server{
		Addr:              "0.0.0.0:9000",
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	http.HandleFunc("/", handler)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed { //would get ErrServerClosed when the server was properly closed
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt) //notify into the channel on SIGINT

	<-stop //wait for stop notice

	//////// shutdown procedure
	// get a server context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	fmt.Printf("Server shutting down")
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++   //lock global variables, as several threads could attempt access
	mu.Unlock()

	fmt.Fprintf(w,"Request Count : %d\n", count)
	fmt.Println("Received request from IP :-> ", r.RemoteAddr)
	fmt.Fprintf(w, "Request path->: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Request method->: %s\n", r.Method)
	fmt.Fprintf(w, "Request headers->:\n")
	for k, v := range r.Header {
		fmt.Fprintf(w, "\t %s: %s\n", k, v)
	}

}
