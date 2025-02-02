package database

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Checkpas(r *http.Request, db *Database) bool {
	var hashPas string
	mailFromVal := r.FormValue("email")

	db.database.QueryRow("select password from users where email = $1", mailFromVal).Scan(&hashPas)

	err := bcrypt.CompareHashAndPassword([]byte(hashPas), []byte(r.FormValue("password")))
	if err != nil {
		log.Printf("invalid pass", err)
		return false
	}

	return true
}
