package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	port := ":8000"

	_, err := connectToDB()
	if err != nil {
		log.Panic("Can't connect to Postgres.")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/events", getAllEvents)

	http.HandleFunc("/events", getAllEvents)
	fmt.Printf("Listening on Port %s\n", port)
	http.ListenAndServe(port, nil)
}

func connectToDB() (*sql.DB, error) {
	dsn := "postgres://ffkmhklc:QtN3se_l72ZTcDgi12K0V1WdY5MMS3i-@flora.db.elephantsql.com/ffkmhklc"
	log.Printf("Connecting to dsn: %s", dsn)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Panicf("Encountered error while connecting to db: %s", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Panicf("Encountered error while pinging to db: %s", err)
		return nil, err
	}
	return db, err
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("alle Events"))
}
