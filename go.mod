module github.com/deemon87/pkg_exporter

go 1.14

require (
	github.com/deemon87/pkg_exporter/collector v0.0.0
	github.com/deemon87/pkg_exporter/info v0.0.0
	github.com/prometheus/client_golang v1.9.0
)

replace (
	github.com/deemon87/pkg_exporter/collector => ./collector
	github.com/deemon87/pkg_exporter/info => ./info
)
