package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	ServerHttp *http.Server
	Mux        *mux.Router
}
