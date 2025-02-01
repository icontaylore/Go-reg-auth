package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"newClient/internal/database"
)

type Server struct {
	ServerHttp *http.Server
	Mux        *mux.Router
}

type ServerDatabase struct {
	dat *database.Database
}
