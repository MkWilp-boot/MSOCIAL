package main

import (
	"html/template"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("public/views/index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
