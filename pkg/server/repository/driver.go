package repository

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Connect ...
func Connect() *sql.DB {

	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
