package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Db        *sql.DB
	closeChan chan bool
)

func Open() {
	closeChan = make(chan bool)
	var err error
	Db, err = sql.Open("mysql",
		"root:@tcp(db:3306)/shortlink")

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
