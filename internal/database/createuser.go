package database

import (
	"fmt"
	"log"
	"net/http"
	"newClient/depended"
	"newClient/internal/validate"
)

func DBHashUser(w http.ResponseWriter, r *http.Request, db *Database) {
	pas, email := r.FormValue("password"), r.FormValue("email")
	hashPasw := validate.HashPas(pas)

	u := &depended.User{
		Email:    email,
		Password: hashPasw,
	}

	CreateUser(w, r, u, db)
}

func CreateUser(w http.ResponseWriter, r *http.Request, user *depended.User, db *Database) {
	email := user.Email
	password := user.Password

	fmt.Println(email, password, "!!!")
	query := `INSERT INTO users (email,password) VALUES ($1, $2)`
	_, err := db.database.Exec(query, email, password)
	if err != nil {
		log.Printf("ошибка записи пользователя")
	} else {
		log.Printf("пользователь создан")
	}

}
