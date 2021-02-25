package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/deemon87/pkg_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var port = flag.Int("port", 19000, "listen port")
	flag.Parse()

	pkg := collector.NewCollector()
	prometheus.MustRegister(pkg)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
