package main

import (
	_ "github.com/lib/pq"
	"newClient/internal/database"
	"newClient/internal/server"
)

func main() {
	s := server.NewServer()
	db := &database.Database{}
	db.ConnectDatabase()

	db.User()
	// Регистрация обработчиков
	s.Handler("/register")
	s.Handler("/login")

	// Запуск сервера
	s.Run()
}
