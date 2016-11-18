package main

import (
	"fmt"
	"io"
	"net/http"
)

// Docs
// https://golang.org/pkg/net/http
// https://golang.org/pkg/io/#Writer

// This is our function we are going to use to handle the request
// All handlers need to accept two arguments
// 1. Is the ResponseWriter interface, we use this to write a reponse back to the client
// 2. Is the Reponse struct which holds useful information about the request headers, method, url etc
func hello(w http.ResponseWriter, r *http.Request) {
	// We use the standard libaries WriteStirng function to write back to the ResponseWriter interface
	// See docs above
	io.WriteString(w, fmt.Sprintf("%s %s", "hello", r.FormValue("name")))
}

func main() {
	// Add ads the function thats going to handle that response
	http.HandleFunc("/", hello)
	// Starts the web server
	// The first argument in this method is the port you want your server to run on
	// The second is a handler. However we have already added this in the line above so we just pass in nil
	http.ListenAndServe(":8000", nil)
}
