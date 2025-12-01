package groupie

import (
    "html/template"
    "net/http"
    "log"
)

// LoadGroup returns group data for templates; replace implementation with your real lookup.
func LoadGroup(id int) interface{} {
    // TODO: lookup and return the real group by id
    return nil
}

func gameHandler(w http.ResponseWriter, r *http.Request) {

    funcMap := template.FuncMap{
        "LoadGroup": LoadGroup, // tu rends ta fonction Go utilisable dans le HTML
    }

    tmpl := template.Must(
        template.New("power4.html").Funcs(funcMap).
            ParseFiles("template/power4.html"),
    )

    // g = ta structure globale ou l’état du jeu
    if err := tmpl.Execute(w, g); err != nil {
        log.Println("❌ Erreur template:", err)
    }
}