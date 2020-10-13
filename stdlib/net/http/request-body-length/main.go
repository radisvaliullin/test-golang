package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {

	// id := 0

	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test ", r.ContentLength)
		payload := struct {
			arg string `json:"arg"`
		}{}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			fmt.Println("payload ", payload)
			fmt.Println("err: ", err)
		}
		fmt.Println("payload ", payload)
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

	// body := strings.NewReader(`{"abc":1234}`)
	body := strings.NewReader(`{}`)
	resp, err := http.Post("http://0.0.0.0:7337/test", "application/json", body)
	if err != nil {
		fmt.Println("post err: ", err)
		return
	}
	resp.Body.Close()
	fmt.Println("done")
}
