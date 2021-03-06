package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// GETContent load all dynamic content to render in template
func GETContent(w http.ResponseWriter, r *http.Request) {
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
		DATE_FORMAT(post_date, '%d/%m/%y') as date_content,
		post_content,
		CASE WHEN post_edited_content = 0 
		THEN "Não editado"
		ELSE "Conteúdo editado" 
		END edicao,
		post_likes
	FROM posts WHERE post_user_id = ?`, 1)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// setting up posts to render main post's page

	var slicePost []post
	for rows.Next() {
		var _post post
		rows.Scan(&_post.Postdate,
			&_post.PostContent,
			&_post.PostEditedContent,
			&_post.PostLikes)
		slicePost = append(slicePost, _post)
	}

	template, err := template.ParseFiles("public/views/home.html")
	if err != nil {
		log.Fatal(err)
	}
	err = template.Execute(w, slicePost)
	if err != nil {
		log.Fatal(err)
	}

}
