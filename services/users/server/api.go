package server

import (
	"github.com/fardinabir/auth-guard/controllers/users"
	"github.com/fardinabir/auth-guard/database"
	"github.com/fardinabir/auth-guard/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
)

func New() (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	err := database.InitDatabase()
	if err != nil {
		log.Println("Postgres connection error", err)
		return nil, err
	}

	err = service.InitRedisClient()
	if err != nil {
		log.Println("Redis connection error", err)
		return nil, err
	}

	userResource := users.NewResource()
	r.Mount("/", userResource.Router())

	return r, nil
}
