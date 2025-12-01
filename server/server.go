package server

import (
	"github.com/furqanrupom/sqlc-crud/config"
	"github.com/furqanrupom/sqlc-crud/handlers"
	"github.com/furqanrupom/sqlc-crud/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	conf *config.Config
	userHandler *handlers.UserHandler
}

func NewServer (conf *config.Config, userHandler *handlers.UserHandler) *Server {
	return &Server{
		conf: conf,
		userHandler: userHandler,
	}
}

func (s *Server) Start() *chi.Mux {
   r := chi.NewRouter()
   r.Use(middleware.Logger)
   r.Use(middleware.Recoverer)
   r.Use(middleware.Heartbeat("/health")) 
   
   r.Route("/api/v1",func(r chi.Router) {
	routes.UserRoutes(r,s.userHandler)
   })
   return r
}