package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	reg := prometheus.NewPedanticRegistry()

	// Histrogram
	temps := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "test_gauge_count",
		Help: "The test gauge cound",
	})
	reg.MustRegister(temps)

	// generate values
	rndSrc := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(rndSrc)
	var i int
	go func() {
		// Simulate some observations.
		for {
			i = rnd.Intn(1000)
			temps.Set(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
			time.Sleep(time.Millisecond * 50)
		}
	}()

	// launch server
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	fmt.Println("start http listen")
	log.Fatal(http.ListenAndServe(":2112", nil))
}
