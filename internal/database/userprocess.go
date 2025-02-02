package database

import (
	"log"
	"net/http"
	"newClient/internal/redirect"
	"newClient/internal/validate"
)

func UserProcessReg(w http.ResponseWriter, r *http.Request, db *Database) {
	if err := validate.Validate(r); err != nil {
		log.Printf("erorr validate", err)
		redirect.RedirectR(w, r, "вводите верные данные")
	}

	if ExistUser(w, r, db) {
		redirect.Redirect(w, r, "пользователь зарегестрирован")
	}

	DBHashUser(w, r, db)
}

func UserProcessLog(w http.ResponseWriter, r *http.Request, db *Database) {
	if err := validate.Validate(r); err != nil {
		log.Printf("erorr validate", err)
		redirect.RedirectR(w, r, "вводите верные данные")
	}

	if ExistUser(w, r, db) {
		Auth(w, r, db)
	} else {
		redirect.RedirectR(w, r, "пользователя не существует")
	}
}
