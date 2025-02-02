package validate

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPas(pas string) string {
	goHash, err := bcrypt.GenerateFromPassword([]byte(pas), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("ошибка в cryptopas..")
	}

	return string(goHash)
}
