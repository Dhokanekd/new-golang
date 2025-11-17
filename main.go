package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()

	fmt.Fprintf(w, "Hello OpenShift! 👋\n")
	fmt.Fprintf(w, "Pod Name: %s\n", host)
	fmt.Fprintf(w, "Request Path: %s\n", r.URL.Path)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // OpenShift expects 8080
	}

	http.HandleFunc("/", handler)

	log.Printf("Server starting on port %s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
