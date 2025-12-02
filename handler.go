package groupie

import (
	"groupie/src"
	"html/template"
	"log"
	"net/http"
)


func LoadGroup(id int) interface{} {
    return nil
}

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
    funcMap := template.FuncMap{
        "LoadGroup": LoadGroup,
    }

    tmpl := template.Must(
        template.New("accueil.html").Funcs(funcMap).
            ParseFiles("template/accueil.html"),
    )
    data := groupie.LoadGroupResum()
    if err := tmpl.Execute(w, data); err != nil {
        log.Println("❌ Erreur template:", err)
    }
}