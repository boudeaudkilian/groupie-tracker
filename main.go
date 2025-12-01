package main

import (
	"log"
	"net/http"
)

func accueilHandler(w http.ResponseWriter, r *http.Request) {
	// Réponse simple pour la route /handler
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("Bonjour et bienvenue sur Groupie Tracker"))
}

func main() {
	//  Association des routes avec leurs fonctions
	http.HandleFunc("/handler", accueilHandler)

	//  Servir les fichiers statiques (CSS, images, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println(" Bonjour et bienvenus sur notre serveur Groupie Tracker il est lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
