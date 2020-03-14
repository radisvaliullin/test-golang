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
	// registry
	reg := prometheus.NewPedanticRegistry()

	// add other types

	// rand
	rndSrc := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(rndSrc)
	var i int

	// // Histrogram
	temps := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "pond_temperature_celsius",
		Help:    "The temperature of the frog pond.", // Sorry, we can't measure how badly it smells.
		Buckets: prometheus.LinearBuckets(20, 5, 5),  // 5 buckets, each 5 centigrade wide.
	})
	reg.MustRegister(temps)

	// Summaries
	tempsSum := prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "pond_temperature_celsius2",
		Help:       "The temperature of the frog pond 2.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})
	reg.MustRegister(tempsSum)

	// gauge
	cpuTemp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
	reg.MustRegister(cpuTemp)

	// Simulate some observations.
	go func() {
		var v float64
		for {
			i = rnd.Intn(1000)
			v = 30 + math.Floor(120*math.Sin(float64(i)*0.1))/10
			temps.Observe(v)
			tempsSum.Observe(v)
			cpuTemp.Set(v)
			time.Sleep(time.Millisecond * 50)
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	fmt.Println("start http listen")
	log.Fatal(http.ListenAndServe(":2112", nil))
}
