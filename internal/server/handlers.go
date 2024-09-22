package server

import (
	"log"
	"net/http"

	"github.com/Frequinzy/roll-some/internal/row"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET: Request for home")
	renderTemplate(w, "home", nil)
}

func putRollSome(w http.ResponseWriter, r *http.Request) {
	log.Printf("PUT: Request for rollSome")

	if err := r.ParseForm(); err != nil {
		log.Print(err.Error())
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	origin := r.FormValue("origin")

	rows, err := row.ParseString(origin)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Unable to parse string", http.StatusBadRequest)
		return
	}

	res := row.RollRows(&rows)

	renderTemplate(w, "rollSome", res)
}
