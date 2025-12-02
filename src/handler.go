package groupie

import (
	"html/template"
	"log"
	"net/http"
)

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
    funcMap := template.FuncMap{
        "LoadGroup": LoadGroup,
    }

    tmpl := template.Must(
        template.New("accueil.html").Funcs(funcMap).
            Funcs(template.FuncMap{
                "mod": func(i, j int) int { return i % j },
                "add1": func(i int) int { return i + 1 },
            }).
            ParseFiles("template/accueil.html"),
    )
    data := LoadGroupResum()
    if err := tmpl.Execute(w, data); err != nil {
        log.Println("❌ Erreur template:", err)
    }
}