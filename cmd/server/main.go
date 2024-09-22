package main

import (
	"log"
	"net/http"

	"github.com/Frequinzy/roll-some/internal/server"
)

func main() {
	mux := server.NewMux()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
