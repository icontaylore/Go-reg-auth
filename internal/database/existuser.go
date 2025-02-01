package database

import (
	"fmt"
	"log"
)

func ExistUser(db *Database) {
	rows, err := db.database.DB.Query("select * from users")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(rows.Columns())
}
