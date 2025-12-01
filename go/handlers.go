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

	rows := make([][]string, 6)
	for rowsIndex := 0; rowsIndex < 6; rowsIndex++ {
		row := make([]string, 7)
		for c := 0; c < 7; c++ {
			row[c] = game.Board[c][rowsIndex]
		}
		rows[rowsIndex] = row
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

// Action du joueur
func PlayHandler(w http.ResponseWriter, r *http.Request) {
	colStr := r.FormValue("col")
	col, _ := strconv.Atoi(colStr)
	if err := game.Play(col - 1); err != nil {
		// On ignore les erreurs
	}

	// Si le joueur gagnant est vide, on continue la partie
	if game.Winner == "" {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/win", http.StatusSeeOther)
	}
}

// Page de fin de partie (win/égalité)
func WinHandler(w http.ResponseWriter, r *http.Request) {
	if game == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	t, _ := template.ParseFiles("static/win.html")
	if t == nil {
		http.Error(w, "Template introuvable", 500)
		return
	}

	data := struct {
		Winner string
		P1Name string
		P2Name string
	}{
		Winner: game.Winner,
		P1Name: game.P1Name,
		P2Name: game.P2Name,
	}

	t.Execute(w, data)
}

// Rematch
func RematchHandler(w http.ResponseWriter, r *http.Request) {
	game.Reset()
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

// Page des règles
func RulesHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("static/regles.html")
	if err != nil {
		http.Error(w, "Erreur template : " + err.Error(), 500)
		return
	}
	template.Execute(w, nil)
}