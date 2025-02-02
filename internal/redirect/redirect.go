package redirect

import (
	"net/http"
	"net/url"
)

func Redirect(w http.ResponseWriter, r *http.Request, er string) {
	encode := url.QueryEscape(er)
	http.Redirect(w, r, "/login?error="+encode, http.StatusSeeOther)
}

func RedirectR(w http.ResponseWriter, r *http.Request, er string) {
	encode := url.QueryEscape(er)
	http.Redirect(w, r, "/register?error="+encode, http.StatusSeeOther)
}
