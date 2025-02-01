package server

import (
	"net/http"
	"net/url"
)


func (s *Server) Redirect(w http.ResponseWriter,r *http.Request, er string) {
	encode := url.QueryEscape(er)
	http.Redirect(w,r,"/login?error="+encode,http.StatusSeeOther)
}