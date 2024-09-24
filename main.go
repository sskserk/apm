package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define a custom metric
var (
	myCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "my_counter_total",
		},
	)

	deTemperatuur = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "serk_laptop_temperatuur",
		},
	)
)

func init() {
	// Register the custom metric with Prometheus
	prometheus.MustRegister(myCounter)
	prometheus.MustRegister(deTemperatuur)
}

func hand(writ http.ResponseWriter, req *http.Request) {
	myCounter.Add(1)

	writ.Write([]byte(myCounter.Desc().String()))
	writ.WriteHeader(http.StatusOK)
}

func temperReader() {
	for {
		deHuidigeTemperatuur := rand.Int()
		fmt.Printf("Current temperature [%d]\n", deHuidigeTemperatuur)
		deTemperatuur.Set(float64(deHuidigeTemperatuur))

		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Increment the counter as an example
	myCounter.Add(10)

	go temperReader()

	// Create a handler for Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/request", hand)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
