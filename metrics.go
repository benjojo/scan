package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	gaugeTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "total",
		Namespace: "scan",
		Subsystem: "ips",
		Help:      "Total IPs found",
	})

	gaugeLatest = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "latest",
		Namespace: "scan",
		Subsystem: "ips",
		Help:      "Latest IPs found",
	})

	gaugeNew = prometheus.NewGauge(prometheus.GaugeOpts{
		Name:      "new",
		Namespace: "scan",
		Subsystem: "ips",
		Help:      "New IPs found",
	})

	gaugeJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "job",
			Namespace: "scan",
			Help:      "Number of IPs found in each each job, with submitted and received times",
		},
		[]string{"id", "submitted", "received"})
)

func init() {
	prometheus.MustRegister(gaugeTotal)
	prometheus.MustRegister(gaugeLatest)
	prometheus.MustRegister(gaugeNew)
	prometheus.MustRegister(gaugeJobs)
}

func metrics() {
	for {
		results, err := resultData("", "", "")
		if err == nil {
			gaugeTotal.Set(float64(results.Total))
			gaugeLatest.Set(float64(results.Latest))
			gaugeNew.Set(float64(results.New))
		}
		time.Sleep(1 * time.Minute)
	}
}
