package groupie

import (
	"html/template"
	"net/http"
	"log"
	)

		
	func LoadGroup(id int) interface{} {
    return nil
}

func gameHandler(w http.ResponseWriter, r *http.Request) {

    funcMap := template.FuncMap{
        "LoadGroup": LoadGroup, // tu rends ta fonction Go utilisable dans le HTML
    }

    tmpl := template.Must(
        template.New("accueil.html").Funcs(funcMap).
            ParseFiles("template/accueil.html"),
    )

    // g = ta structure globale ou l’état du jeu
    if err := tmpl.Execute(w, g); err != nil {
        log.Println("❌ Erreur template:", err)
    }
}
