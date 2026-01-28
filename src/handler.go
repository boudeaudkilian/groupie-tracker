package groupie

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type RequestData struct {
	Input string `json:"input"`
}

type ResponseData struct {
	Result string `json:"result"`
}

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"LoadGroup": LoadGroup,
	}

	tmpl := template.Must(
		template.New("accueil.html").Funcs(funcMap).
			Funcs(template.FuncMap{
				"mod":  func(i, j int) int { return i % j },
				"add1": func(i int) int { return i + 1 },
			}).
			ParseFiles("template/accueil.html"),
	)

	data := LoadGroupResum()
	if err := tmpl.Execute(w, data); err != nil {
		log.Println("❌ Erreur template:", err)
	}
}

func PageGroupHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"LoadGroup": LoadGroup,
	}

	tmpl := template.Must(
		template.New("grppage.html").Funcs(funcMap).
			ParseFiles("template/grppage.html"),
	)

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	grp := LoadGroup(id)
	if grp == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := tmpl.Execute(w, grp); err != nil {
		log.Println("❌ Erreur template:", err)
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.New("searchbar.html").ParseFiles("template/searchbar.html"),
	)

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println("❌ Erreur template searchbar:", err)
	}
}

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}
	
	fmt.Println(data.Input)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}