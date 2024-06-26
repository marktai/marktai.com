package db

import (
	"database/sql"
	_ "github.com/lib/pq"

	"log"
)

var (
	Db        *sql.DB
	closeChan chan bool
)

func Open() {
	closeChan = make(chan bool)
	var err error
	Db, err = sql.Open("postgres",
		"host=127.0.0.1 port=5432 user=postgres password=postgres dbname=shortlink sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	go func() {
		<-closeChan
		Db.Close()
	}()
}

func Close() {
	closeChan <- true
}
