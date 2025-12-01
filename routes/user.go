package routes

import (
	"github.com/furqanrupom/sqlc-crud/handlers"
	"github.com/go-chi/chi"
)

func UserRoutes(r chi.Router, h *handlers.UserHandler){
   r.Post("/users",h.CreateUser)
   r.Get("/users",h.GetUsers)
   r.Get("/users/{id}",h.GetUser)
}