package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func NewServer() *Server {
	return &Server{
		Mux: mux.NewRouter(),
	}
}

func (s *Server) Run() *http.Server {
	serv := &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		Handler:      s.Mux,
	}

	fmt.Print("[SERVER START]")
	if err := serv.ListenAndServe(); err != nil {
		log.Printf("TROUBLE LIST")
	}
	return serv
}
