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

	// Регистрация обработчиков
	s.Handler("/register", db)
	s.Handler("/login", db)

	// Запуск сервера
	s.Run()
}
