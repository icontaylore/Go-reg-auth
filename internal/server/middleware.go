package server

import (
	"fmt"
	"net/http"
)


func LogMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL)

		next(w,r)
	}
}