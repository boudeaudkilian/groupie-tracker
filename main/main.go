package main

import (
	"groupie"
	"log"
	"net/http"
)

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("Bonjour et bienvenue sur Groupie Tracker"))
}

func main() {
	http.HandleFunc("/", groupie.AccueilHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println(" Bonjour et bienvenus sur notre serveur Groupie Tracker il est lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
