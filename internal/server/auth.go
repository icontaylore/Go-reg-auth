package server

import (
	"fmt"
	"net/http"
	"newClient/internal/htmlpars"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	errorik := r.URL.Query().Get("error")
	fmt.Print(errorik)

	if r.Method == http.MethodGet {
		htmlpars.ParsingLog(w, r, errorik)
	}
	if r.Method == http.MethodPost {
		s.Redirect(w, r, "")
	}

}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	errorik := r.URL.Query().Get("error")
	fmt.Print(errorik)

	if r.Method == http.MethodGet {
		htmlpars.ParsingReg(w, r, errorik)
	}

	if r.Method == http.MethodPost {

	}
}
