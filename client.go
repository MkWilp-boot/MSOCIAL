package main

import (
	"fmt"
	"net/http"
)

type post struct {
	Postdate          string
	PostContent       string
	PostEditedContent int
	PostLikes         int
}

// Page refers to all data to be sent to the client
type Page struct {
	posts []post
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		GETContent(w, r)
	case r.Method == "POST":
		InsertPOSTContent(w, r)
	default:
		fmt.Fprint(w, "<h1>Error, method now allowed</h1>")
	}
}
