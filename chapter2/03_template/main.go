package main

import (
	"html/template"
	"net/http"
	"strings"
	"log"
)

var (
	helloTemplate *template.Template
)

func init() {
	var err error
	helloTemplate, err = template.New("hello.tmpl").ParseFiles("./hello.tmpl")
	if err != nil {
		panic(err)
	}
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	pathParams := strings.Split(req.URL.Path, "/")
	params := struct {
		Name string
	}{
		pathParams[2],
	}
	err := helloTemplate.Execute(w, params)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/hello/", HelloHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
