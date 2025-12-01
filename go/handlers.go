package groupie

import (
	"html/template"
	"net/http"
	"strconv"
)

var game *Game

//Page accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		template, err := template.ParseFiles("static/start.html")
		if err != nil {
			http.Error(w, "Erreur template : "+err.Error(), 500)
			return
		}
		template.Execute(w, nil)
	case http.MethodPost:
		p1 := r.FormValue("p1")
		p2 := r.FormValue("p2")
		if p1 == "" {
			p1 = "Joueur 1"
		}
		if p2 == "" {
			p2 = "Joueur 2"
		}
		game = InitGame(p1, p2)
		http.Redirect(w, r, "/game", http.StatusSeeOther)
	}
}

// Page du jeu
func GameHandler(w http.ResponseWriter, r *http.Request) {
	if game == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	t, _ := template.ParseFiles("static/index.html")
	data := struct {
		Rows   [][]string
		Player string
		Winner string
		P1Name string
		P2Name string
	}{
		Rows:   rows,
		Player: game.Player,
		Winner: game.Winner,
		P1Name: game.P1Name,
		P2Name: game.P2Name,
	}
	t.Execute(w, data)
}