package validate

import (
	"github.com/go-playground/validator"
	"log"
	"net/http"
)

type User struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func Validate(r *http.Request) error {
	user := &User{Email: r.FormValue("email"), Password: r.FormValue("password")}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		log.Printf("validation error, [PAS or EMAIL]")
		return err
	}

	return nil
}
