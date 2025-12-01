package main

import (
	"log"
	"net/http"
	"groupie"
)

func main() {
	//  Association des routes avec leurs fonctions
	http.HandleFunc("/handler", accueilHandler)

	//  Servir les fichiers statiques (CSS, images, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println(" Bonjour et bienvenus sur notre serveur Groupie Tracker il est lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
