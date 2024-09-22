package server

import (
	"net/http"
)

func NewMux() *http.ServeMux {
	loadTemplates()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", getHome)
	mux.HandleFunc("POST /rollSome", putRollSome)

	return mux
}
