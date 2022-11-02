package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func InitDB() (Database, error) {
	db := Database{}
	var host = "localhost"
	var port = 5432
	var username = os.Getenv("POSTGRES_USER")
	var password = os.Getenv("POSTGRES_PASSWORD")
	var database = os.Getenv("POSTGRES_DB")
	dbStringConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)

	conn, err := sql.Open("postgres", dbStringConnection)

	if err != nil {
		log.Println("Database connection error", err)
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()

	if err != nil {
		log.Println("Database connection error", err)
		return db, err
	}

	log.Println("Database connection established")
	return db, nil

}
