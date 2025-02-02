package server

import (
	"fmt"
	"net/http"
	"newClient/internal/database"
	"newClient/internal/htmlpars"
	"newClient/internal/redirect"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request, db *database.Database) {
	errorik := r.URL.Query().Get("error")
	fmt.Print(errorik)

	if r.Method == http.MethodGet {
		htmlpars.ParsingLog(w, r, errorik)
	}
	if r.Method == http.MethodPost {
		redirect.Redirect(w, r, "")
	}

}

func (s *Server) Register(w http.ResponseWriter, r *http.Request, db *database.Database) {
	errorik := r.URL.Query().Get("error")
	fmt.Print(errorik)

	if r.Method == http.MethodGet {
		htmlpars.ParsingReg(w, r, errorik)
	}

	if r.Method == http.MethodPost {
		database.UserProcessReg(w, r, db)
	}
}
