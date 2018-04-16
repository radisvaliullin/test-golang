package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Test context for net.http package
func main() {

	// run test servers
	// servers accept our non blocking requests
	startHTTPServ()

	// init context variable
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	// setup context with timeout, after timeout expiration all request will abort
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*6)
	_ = cancel

	// non blocking resp api req (to our test http server)
	res, err := restReq(ctx)
	if err != nil {
		log.Fatal("call restReq: err - ", err)
	}
	log.Printf("call restReq res: %v", string(res))
}

// rest api req
func restReq(ctx context.Context) ([]byte, error) {

	// build request
	req, err := http.NewRequest("GET", "http://localhost:7373/longlived", nil)
	if err != nil {
		log.Println("restRec: new req err - ", err)
		return nil, err
	}
	req = req.WithContext(ctx)

	// build client
	cln := &http.Client{}

	// do req
	resp, err := cln.Do(req)
	if err != nil {
		log.Print("restReq: do req, err - ", err)
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("restReq: read resp body, err - ", err)
		return nil, err
	}

	return respBytes, nil
}

// servers

// simple http rest api handle server
func startHTTPServ() {

	mux := http.NewServeMux()

	// long-lived handler
	mux.HandleFunc("/longlived", func(w http.ResponseWriter, r *http.Request) {
		// long lived, delay
		time.Sleep(time.Second * 5)
		w.Write([]byte("Long lived request"))
	})

	// run server listener
	go func() {
		if err := http.ListenAndServe(":7373", mux); err != nil {
			log.Fatal("http server down: err - ", err)
		}
	}()

}
