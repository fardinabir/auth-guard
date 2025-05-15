package users

import (
	"github.com/fardinabir/auth-guard/db/repos"
	"github.com/fardinabir/auth-guard/service"
	"github.com/go-chi/chi/v5"
)

type UserResource struct {
	Users       UserStore
	TokenClient *service.TokenClient
}

func NewResource() *UserResource {
	userStore := repos.NewUserStore()
	tokenClient := service.NewTokenClient()
	return &UserResource{Users: userStore, TokenClient: tokenClient}
}

func (rs *UserResource) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use()

	r.Mount("/users", rs.userRouter())
	r.Post("/login", rs.Login)
	r.Get("/", rs.HomePage)
	return r
}

func (rs *UserResource) userRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(service.WardenStampChecker)
	r.Use(service.UserNameChecker)
	r.Get("/{id:[0-9]+}", rs.ReadUser)
	r.Get("/", rs.ReadUsers)
	r.Post("/", rs.CreateUser)
	r.Patch("/{id:[0-9]+}", rs.UpdateUser)
	r.Delete("/{id:[0-9]+}", rs.DeleteUser)

	return r
}
