package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

//
var cntMx = sync.Mutex{}
var connCnt = 0

func main() {
	// monitor http server
	monSrvAddr := "0.0.0.0:7373"
	go func() {
		http.HandleFunc("/monitor", func(w http.ResponseWriter, r *http.Request) {
			// count value
			cnt := -1
			cntMx.Lock()
			cnt = connCnt
			cntMx.Unlock()

			// response json
			mon := struct {
				ConnCount int `json:"connCount"`
			}{
				ConnCount: cnt,
			}
			monJSON, err := json.Marshal(&mon)
			if err != nil {
				log.Printf("monitor json marshal err - %v", err)
				return
			}

			// write response
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(monJSON)
			if err != nil {
				log.Printf("monitor json write err - %v", err)
				return
			}
		})
		if err := http.ListenAndServe(monSrvAddr, nil); err != nil {
			log.Printf("monitor server err - %v", err)
		}
	}()

	// connect to db
	dataSource := "postgres://postgres:postgres@0.0.0.0:5432/test?sslmode=disable"

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatalf("open db err, %v", err)
	}
	defer db.Close()
	log.Printf("db open, with params - %v", dataSource)

	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping err, %v", err)
	}
	log.Print("db ping, ok")

	//
	// test limit of connection, test connection pool
	//

	// parallel long life reqeusts
	go func() {
		db.SetMaxOpenConns(30)
		wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go longLifeDBReq(i, wg, db)
		}
		wg.Wait()
		log.Printf("long life parallel request end")
	}()

	// prepare
	// handle ctrl-c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Printf("ctrl-c pressed")
	err = db.Close()
	if err != nil {
		log.Printf("db close err - %v", err)
	}
}

func longLifeDBReq(reqID int, wg *sync.WaitGroup, db *sql.DB) {
	defer wg.Done()

	life := 10
	prefMsg := "long life db request"
	now := time.Now()
	log.Printf("%v: start, id - %v", prefMsg, reqID)

	cntMx.Lock()
	connCnt++
	cntMx.Unlock()
	res, err := db.Exec("SELECT pg_sleep($1);", life)
	if err != nil {
		log.Printf("%v: request, id - %v, err - %v", prefMsg, reqID, err)
		return
	}
	cntMx.Lock()
	connCnt--
	cntMx.Unlock()
	_ = res

	log.Printf("%v: end, id - %v, perform time - %v", prefMsg, reqID, time.Since(now).Seconds())
}
