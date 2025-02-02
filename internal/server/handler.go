package server

import (
	"net/http"
	"newClient/internal/database"
)

func (s *Server) Handler(path string, db *database.Database) {
	switch path {
	case "/register":
		s.Mux.HandleFunc(path, LogMiddleWare(func(w http.ResponseWriter, r *http.Request) {
			s.Register(w, r, db)
		}))
	case "/login":
		s.Mux.HandleFunc(path, LogMiddleWare(func(w http.ResponseWriter, r *http.Request) {
			s.Login(w, r, db)
		}))
	}
}
