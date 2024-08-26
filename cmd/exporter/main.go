package main

import (
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	collector "blade_status_exporter/cmd/collector"
)

func main() {
	config, err := GetConf("./files/config.yaml")
	if err != nil {
		return
	}

	chassisCollector := collector.NewChassisCollector(config.ChassisList, config.Username, config.Password)
	registry := prometheus.NewRegistry()

	registry.MustRegister(chassisCollector)

	promHandler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	})

	http.Handle("/metrics", promHandler)
	log.Fatal(http.ListenAndServe(":9101", nil))
}
