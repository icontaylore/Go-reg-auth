package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	database *sqlx.DB
}

func UsingDb() *Database {
	return &Database{}
}

func (d *Database) ConnectDatabase() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=1")
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Ошибка при пинге базы данных: %v", err)
	}

	fmt.Printf("[DATABASE START]\n")
	d.database = db
	return d.database
}
