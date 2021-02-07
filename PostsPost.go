package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

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
		log.Fatalf("Statement prepare failure, full stacetrack %s\n", err)
	}
	stmt.Exec(r.Form.Get("AreaComment"), 1, 0, 0)
	GETContent(w, r)
}
