package database

import (
	"fmt"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request, db *Database) {
	if Checkpas(r, db) {
		fmt.Println("[PASS TRUE]")

	}

}
