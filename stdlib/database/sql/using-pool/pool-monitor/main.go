package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {

	srvAdd := "0.0.0.0:7373"
	rURL := url.URL{
		Scheme: "http",
		Host:   srvAdd,
		Path:   "monitor",
	}

	for {
		toSrvGet(rURL.String())
		time.Sleep(time.Second)
	}
}

func toSrvGet(rURL string) {
	res, err := http.Get(rURL)
	if err != nil {
		log.Printf("request err - %v", err)
		return
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("read body err - %v", err)
		return
	}

	mon := struct {
		ConnCount int `json:"connCount"`
	}{}
	err = json.Unmarshal(b, &mon)
	if err != nil {
		log.Printf("body unmarshal err - %v", err)
		return
	}

	log.Printf("monitor - %+v", mon)
}
