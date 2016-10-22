package main

import (
	"io"
	"net/http"
	"log"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
