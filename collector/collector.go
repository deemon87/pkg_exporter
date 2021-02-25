package collector

import (
	"github.com/deemon87/pkg_exporter/info"

	"github.com/prometheus/client_golang/prometheus"
)

//Collector type
type Collector struct {
	PkgInfo *prometheus.Desc
}

// NewCollector func
func NewCollector() *Collector {
	return &Collector{
		PkgInfo: prometheus.NewDesc("pkg_info",
			"Installed packeges",
			[]string{"name", "version", "desired_action", "package_status", "error_flags"}, nil,
		),
	}
}

// Describe method
func (collector *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.PkgInfo
}

// Collect method
func (collector *Collector) Collect(ch chan<- prometheus.Metric) {
	packages := info.GetPkgInfo()
	for _, pkg := range *packages {
		ch <- prometheus.MustNewConstMetric(collector.PkgInfo, prometheus.GaugeValue, 0,
			pkg.Name, pkg.Version, pkg.DesiredAction, pkg.PackageStatus, pkg.ErrorFlags)
	}
}
