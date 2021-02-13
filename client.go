package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

func initLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	username := r.Form.Get("username")
	usPpasswd := r.Form.Get("password")

	ip, err := sess.GetIP(r)

	ok := sess.SetSession(ip, username, usPpasswd)

	if ok != nil {
		fmt.Fprint(w, ok)
	} else {
		fmt.Fprintf(w,
			"Welcome user! Today is %s",
			time.Now().Format("02/01/2006 03:04:05"))
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var continueNavigation bool

	ip, err := sess.GetIP(r)
	if err != nil {
		panic(err)
	}
	logged := sess.Session(ip)
	if logged {
		logged = sess.CheckSession(ip)
	}
	if !logged {
		continueNavigation = false
	} else {
		continueNavigation = true
	}

	switch {
	case r.Method == "GET" && continueNavigation:
		GETContent(w, r)
	case r.Method == "POST" && continueNavigation:
		InsertPOSTContent(w, r)
	case !continueNavigation:
		loginHandler(w, r)
	default:
		fmt.Fprint(w, "<h1>Error, method now allowed</h1>")
	}
}
