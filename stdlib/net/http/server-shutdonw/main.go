package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {

	id := 0

	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		id := id + 1
		fmt.Println("ID: ", id)
		time.Sleep(time.Second * 600)
		fmt.Println("end ID: ", id)
		_, err := w.Write([]byte(fmt.Sprintf("done. ID %v", id)))
		if err != nil {
			fmt.Println("err ID ", id)
		}
	})

	srv := http.Server{
		Addr:         "0.0.0.0:7337",
		Handler:      mux,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println("server down: ", err)
		}
	}()

	time.Sleep(time.Second)

	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get("http://0.0.0.0:7337/test")
			if err != nil {
				fmt.Println("get err ", err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("body read err ", err)
			}
			fmt.Println("body ", string(body))
		}()
	}

	time.Sleep(time.Second * 3)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("shutdonw start")
		if err := srv.Shutdown(context.Background()); err != nil {
			fmt.Println("shutdown err ", err)
			return
		}
		fmt.Println("shutdonw end")
	}()

	wg.Wait()
	fmt.Println("done")
}
