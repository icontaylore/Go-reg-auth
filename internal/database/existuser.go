package database

import (
	"log"
	"net/http"
)

func ExistUser(w http.ResponseWriter, r *http.Request, db *Database) bool {
	var count bool
	err := db.database.Get(&count, "SELECT COUNT(*) from users WHERE email = $1", r.FormValue("email"))
	if err != nil {
		log.Println(err)
	}

	return count
}
