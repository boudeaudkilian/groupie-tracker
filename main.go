package main

import (
	"html/template"
	"log"
	"net/http"
	//"groupie"
)

func main() {
	//  Association des routes avec leurs fonctions
	http.HandleFunc("/", accueilHandler)

	//  Servir les fichiers statiques (CSS, images, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println(" Bonjour et bienvenus sur notre serveur Groupie Tracker il est lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func accueilHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/accueil.html"))
	tmpl.Execute(w, nil)
}
