package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var Version string
var Buildtime string

func main() {
	go func() {
		for range time.NewTicker(time.Second).C {
			log.Printf("This is a heartbeat..." + time.Now().String())
		}
	}()
	// use PORT environment variable, or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// register hello function to handle all requests
	server := http.NewServeMux()
	server.HandleFunc("/", hello)

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request: %s", r.URL.Path)
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "Version: %s\n", Version)
	fmt.Fprintf(w, "Build time: %s\n", Buildtime)
	fmt.Fprintf(w, "D2iQ CI/CD!\n")
}
