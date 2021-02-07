package main

import (
	"fmt"
	"net/http"

	"github.com/MkWilp-boot/sess"
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
	ip, err := sess.GetIP(r)
	logged := sess.Session(ip)
	if logged {
		sess.SetSession(ip)
	}

	switch {
	case r.Method == "GET" && err == nil:
		//sess.SetSession(ip)
		GETContent(w, r)
	case r.Method == "POST" && err == nil:
		InsertPOSTContent(w, r)
	case err != nil:
		fmt.Fprintf(w, "<h1>You are not allowed to access this website</h1>")
	default:
		fmt.Fprint(w, "<h1>Error, method now allowed</h1>")
	}
}
