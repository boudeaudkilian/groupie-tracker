package main

import (
	groupie "groupie/src"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", groupie.AccueilHandler)
	http.HandleFunc("/search", groupie.SearchHandler)
	http.HandleFunc("/group", groupie.PageGroupHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Bonjour et bienvenus sur notre serveur Groupie Tracker: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
