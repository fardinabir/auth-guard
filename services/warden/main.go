package main

import (
	"log"
	"net/http"

	"github.com/fardinabir/auth-guard/services/warden/api"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	tokenHandler := api.NewTokenHandler()

	r.Post("/tokens", tokenHandler.GenerateToken)
	r.Get("/tokens/validate", tokenHandler.ValidateToken)

	log.Fatal(http.ListenAndServe(":8081", r))
}
