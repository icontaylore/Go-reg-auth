package htmlpars

import (
	"log"
	"net/http"
	"text/template"
)



func ParsingReg(w http.ResponseWriter, r *http.Request, errorik string) {
	tmpl,err := template.ParseFiles("../templates/register.html")
	if err != nil {
		log.Printf("error parsing html")
	}

	tmpl.Execute(w, struct {
		Error string
	}{Error: errorik})
}

func ParsingLog(w http.ResponseWriter, r *http.Request, errorik string) {
	tmpl,err := template.ParseFiles("../templates/login.html")
	if err != nil {
		log.Printf("error parsing html")
	}

	tmpl.Execute(w, struct {
		Error string
	}{Error: errorik})
}

