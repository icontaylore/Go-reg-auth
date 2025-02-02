package server

import (
	"log"
	"net/http"
	"newClient/internal/database"
	"newClient/internal/htmlpars"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request, db *database.Database) {
	errorik := r.URL.Query().Get("error")
	log.Printf("ошибка - ", errorik)

	if r.Method == http.MethodGet {
		htmlpars.ParsingLog(w, r, errorik)
	}
	if r.Method == http.MethodPost {
		database.UserProcessLog(w, r, db)
	}

}

func (s *Server) Register(w http.ResponseWriter, r *http.Request, db *database.Database) {
	errorik := r.URL.Query().Get("error")
	log.Printf("ошибка - ", errorik)

	if r.Method == http.MethodGet {
		htmlpars.ParsingReg(w, r, errorik)
	}

	if r.Method == http.MethodPost {
		database.UserProcessReg(w, r, db)
	}
}
