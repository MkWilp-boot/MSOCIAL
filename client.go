package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type post struct {
	Postdate          string
	PostContent       string
	PostEditedContent int8
	PostLikes         int
}

// Page refers to all data to be sent to the client
type Page struct {
	title string
	posts []post
}

// InsertPOSTContent is designed for uploading posts
func InsertPOSTContent(w http.ResponseWriter, r *http.Request) {
	// DataBase conn and post searching

	db, err := sql.Open("mysql", "joaoR:Joao_846515_AX@/MSOCIAL")
	if err != nil {
		ServerErrorTemplate, err := template.ParseFiles("public/internalServerError.html")
		if err != nil {
			fmt.Fprint(w, "<h1>Server under maitence, come back later</h1>")
			log.Fatalf("Error rendering 500+ template, stace track: %s", err)
		}
		ServerErrorTemplate.Execute(w, nil)
		log.Fatalf("stacetarck: %s", err)
	}
	defer db.Close()

	err = r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare(`INSERT INTO posts(
			post_content,
			post_user_id,
			post_edited_content,
			post_likes
		)
		VALUES(?, ?, ?, ?)`)

	if err != nil {
		log.Fatalf("Statement prepare failure, full stace track %s\n", err)
	}
	stmt.Exec(r.Form.Get("AreaComment"), 1, 0, 0)

	fmt.Fprintf(w, "<h1>Post sent: %s</h1>", r.Form.Get("AreaComment"))

	template, err := template.ParseFiles("public/home.html")
	if err != nil {
		log.Fatal(err)
	}
	err = template.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// GETContent load all dynamic content to render in template
func GETContent(w http.ResponseWriter, r *http.Request) {
	var rec int

	db, err := sql.Open("mysql", "joaoR:Joao_846515_AX@/MSOCIAL")
	if err != nil {
		ServerErrorTemplate, err := template.ParseFiles("public/internalServerError.html")
		if err != nil {
			fmt.Fprint(w, "<h1>Server under maitence, come back later</h1>")
			log.Fatalf("Error rendering 500+ template, stace track: %s", err)
		}
		ServerErrorTemplate.Execute(w, nil)
		log.Fatalf("stacetarck: %s", err)
	}
	defer db.Close()
	rows, err := db.Query(`SELECT 
		post_date,
		post_content,
		post_edited_content,
		post_likes
	FROM posts WHERE post_user_id = ?`, 1)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	db.QueryRow(`select 
		count(*) as cnt 
		FROM posts 
	WHERE post_user_id = ? LIMIT 1`, 1).Scan(&rec)

	// setting up posts to render main post's page
	slicePost := make([]post, rec)
	for rows.Next() {
		var _post post
		rows.Scan(&_post.Postdate,
			&_post.PostContent,
			&_post.PostEditedContent,
			&_post.PostLikes)
		slicePost = append(slicePost, _post)
	}

	newPage := Page{
		title: "Main Page",
		posts: slicePost,
	}

	template, err := template.ParseFiles("public/home.html")
	if err != nil {
		log.Fatal(err)
	}
	err = template.Execute(w, newPage)
	if err != nil {
		log.Fatal(err)
	}

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
