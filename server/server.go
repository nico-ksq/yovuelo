package server

import (
	"database/sql"

	"github.com/gorilla/mux"
	"yovuelo/api/user_registration"
)

// Server estructura del servidor con el enrutador
type Server struct {
	Router *mux.Router

	user user_registration.UserRegistration
}

// NewServer crea una nueva instancia del servidor
func NewServer(db *sql.DB) *Server {
	return &Server{
		Router: mux.NewRouter(),

		user: user_registration.New(db),
	}
}

// SetupRoutes configura las rutas del servidor
func (s *Server) SetupRoutes() {
	s.Router.HandleFunc("/register", s.UserHandler).Methods("POST")
}
