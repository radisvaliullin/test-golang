package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	envDBName = "APP_DB_NAME"
	envDBUser = "APP_DB_USER"
	envDBPWD  = "APP_DB_PWD"
)

func main() {

	connStr := fmt.Sprintf("user=%v password=%v dbname=%v",
		os.Getenv(envDBUser), os.Getenv(envDBPWD), os.Getenv(envDBName))

	// simple
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ping
	err = db.Ping()
	if err != nil {
		log.Fatal("ping err ", err)
	}
	log.Println("ping OK")

	// create
	sqlTmpl := "CREATE TABLE temp_table(id integer PRIMARY KEY, name text);"
	res, err := db.Exec(sqlTmpl)
	if err != nil {
		log.Fatal("db exec, err ", err)
	}
	log.Printf("db exec res - %+v", res)

	// drop table
	sqlTmpl = "DROP TABLE temp_table;"
	res, err = db.Exec(sqlTmpl)
	if err != nil {
		log.Fatal("db exec, err ", err)
	}
	log.Printf("db exec res - %+v", res)
}
