package main

import (
	"context"
	"database/sql"
	"log"
)

func main() {

	db, err := sql.Open("", "")
	if err != nil {
		log.Fatal("db open ", err)
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		log.Fatal("begin tx ", err)
	}

	_, err = tx.Exec("")
	if err != nil {
		tx.Rollback()
		log.Fatal("exec ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("commit ", err)
	}

}
