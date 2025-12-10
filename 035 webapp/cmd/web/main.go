package main

import (
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
	"time"
)

const webPort = "80"

func main() {

	// database connection
	db := initDB()
	db.Ping()
	// sessions

	// channels

	// waitgroup

	// set up the application config

	// set up mail

	// listen for web connections

}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("failed to connect to database")
	}
	return conn
}

func connectToDB() *sql.DB {
	counts := 0
	dsn := os.Getenv("DSN")

	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet ready")
		} else {
			log.Println("connected to postgres")
			return conn
		}

		if counts > 10 {
			return nil
		}
		log.Println("backing off for 1 second")
		time.Sleep(1 * time.Second)
		counts++

		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
