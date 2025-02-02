package database

import (
	"net/http"
	"newClient/internal/redirect"
)

func UserProcessReg(w http.ResponseWriter, r *http.Request, db *Database) {
	if ExistUser(w, r, db) {
		redirect.Redirect(w, r, "пользователь зарегестрирован")
	}

	CreateUser(w, r, db)
}
