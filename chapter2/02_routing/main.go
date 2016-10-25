package main

import (
	"io"
	"net/http"
	"log"
)

func NotFoundHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "request path not found\n")
}

func GreetingHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hi, forks!\n")
}

func ByeHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bye bye!\n")
}

func main() {
	http.HandleFunc("/", NotFoundHandler)
	http.HandleFunc("/greeting/", GreetingHandler)
	http.HandleFunc("/greeting/bye", ByeHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

