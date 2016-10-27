package main

import (
	"html/template"
	"net/http"
	"log"
)

var (
	chatTemplate *template.Template
	posts []Post
)

type Post struct {
	Name string
	Message string
}

func init() {
	var err error
	chatTemplate, err = template.New("chat.tmpl").ParseFiles("./chat.tmpl")
	if err != nil {
		panic(err)
	}
}

func ChatHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		newPost := Post {
			Name: req.PostFormValue("name"),
			Message: req.PostFormValue("message"),
		}
		posts = append(posts, newPost)
	}

	params := map[string]interface{} {
		"Posts": posts,
	}
	err := chatTemplate.Execute(w, params)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", ChatHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
