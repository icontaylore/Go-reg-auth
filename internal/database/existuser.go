package database

import (
	"fmt"
	"log"
	"net/http"
)

func ExistUser(w http.ResponseWriter, r *http.Request, db *Database) bool {
	var count bool
	err := db.database.Get(&count, "SELECT COUNT(*) from users WHERE email = $1", r.FormValue("email"))
	if err != nil {
		log.Println(err)
	}
	fmt.Println(count, r.FormValue("email"))
	return count
}
