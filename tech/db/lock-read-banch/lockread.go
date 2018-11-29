package main

import (
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/lib/pq"
)

func main() {

	// connect to test docker db
	dataSource := "postgres://postgres:postgres@0.0.0.0:5454/postgres?sslmode=disable"

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

	// // create test datas
	// q := `
	// 	CREATE TABLE test (
	// 		id serial,
	// 		updated_at timestamptz not null default current_timestamp,
	// 		primary key (id)
	// 	);
	// `
	// _, err = db.Exec(q)
	// if err != nil {
	// 	log.Fatalf("create table request err, %v", err)
	// }

	// // set data
	// q := `
	// 	INSERT INTO test (updated_at) VALUES (DEFAULT);
	// `
	// for i := 0; i < 10; i++ {
	// 	_, err = db.Exec(q)
	// 	if err != nil {
	// 		log.Fatalf("insert request err, %v", err)
	// 	}
	// }

	wg := sync.WaitGroup{}
	f := func(gid int, offset int) {
		defer wg.Done()
		err = exclusiveLockReadRowBatch(gid, db, 5, offset)
		if err != nil {
			// log.Printf("gid - %v, exclusive lock read err, %v", gid, err)
			// return
			if !strings.HasPrefix(err.Error(), "pq: could not obtain lock on row in relation") {
				log.Printf("gid - %v, exclusive lock read err, %v", gid, err)
				return
			}
			// time.Sleep(time.Millisecond * 2000)
			log.Printf("git - %v, exclusive lock read err, %v, try again", gid, err)
			err = exclusiveLockReadRowBatch(gid, db, 5, offset+5)
			if err != nil {
				log.Printf("gid - %v, exclusive lock read 2 err, %v", gid, err)
				return
			}
		}
	}

	wg.Add(1)
	go f(0, 0)
	wg.Add(1)
	go f(1, 0)

	wg.Wait()
	log.Print("result, ok")
}

func exclusiveLockReadRowBatch(gid int, db *sql.DB, limit int, offset int) (err error) {
	var tx *sql.Tx
	defer func() {
		if err != nil {
			if tx != nil {
				tx.Rollback()
				log.Printf("gid - %v, rollback", gid)
			}
		}
	}()

	// TX
	tx, err = db.Begin()
	if err != nil {
		log.Printf("gid - %v, exclusive tx begin err, %v", gid, err)
		return err
	}

	// select batch
	bIDs := []int64{}
	q := `
		SELECT id FROM test ORDER BY updated_at ASC LIMIT $1 OFFSET $2;
	`
	r, err := tx.Query(q, limit, offset)
	if err != nil {
		log.Printf("gid - %v, query batch ids err, %v", gid, err)
		return err
	}
	defer r.Close()
	for r.Next() {
		var id int64
		if err := r.Scan(&id); err != nil {
			log.Printf("gid - %v, query batch ids parse err, %v", gid, err)
			return err
		}
		bIDs = append(bIDs, id)
	}
	if err = r.Err(); err != nil {
		log.Printf("gid - %v, query batch ids row err, %v", gid, err)
		return err
	}

	// lock select for update
	q = `
		SELECT id, updated_at FROM test WHERE id = ANY($1) FOR update NOWAIT;
	`
	r, err = tx.Query(q, pq.Array(bIDs))
	if err != nil {
		log.Printf("gid - %v, exclusive query err, %v", gid, err)
		return err
	}
	defer r.Close()
	for r.Next() {
		var id int64
		var tm time.Time
		if err := r.Scan(&id, &tm); err != nil {
			log.Printf("gid - %v, exlusive query row parse err, %v", gid, err)
			return err
		}
		log.Printf("gid - %v, row id - %v, tm - %v", gid, id, tm)
	}
	if err = r.Err(); err != nil {
		log.Printf("gid - %v, exclusive query row err, %v", gid, err)
		return err
	}

	_, err = tx.Exec("SELECT pg_sleep(3);")
	if err != nil {
		log.Printf("gid - %v, exclusive query pg sleep err, %v", gid, err)
		return err
	}

	tx.Commit()
	log.Printf("gid - %v, commited", gid)
	return nil
}
