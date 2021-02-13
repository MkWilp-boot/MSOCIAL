package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", initLogin)

	http.Handle("/public/",
		http.StripPrefix("/public/",
			http.FileServer(http.Dir("public"))))

	log.Println("Executando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
