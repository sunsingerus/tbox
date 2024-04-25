package collectors

import (
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
)

// BuildInfoCollector is a collector for build information metrics
type BuildInfoCollector struct {
	buildInfoDesc *prometheus.Desc
}

// NewBuildInfoCollector creates new BuildInfoCollector
func NewBuildInfoCollector(name, version, tag, sha, built string) *BuildInfoCollector {
	return &BuildInfoCollector{
		buildInfoDesc: prometheus.NewDesc(
			"build_info",
			"Information about the build.",
			nil,
			prometheus.Labels{
				"name":            name,
				"version":         version,
				"tag":             tag,
				"git_sha":         sha,
				"built_at":        built,
				"runtime_version": runtime.Version(),
			},
		),
	}
}

// Describe returns all descriptions of the collector.
func (c *BuildInfoCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.buildInfoDesc
}

// Collect returns the current state of all metrics of the collector.
func (c *BuildInfoCollector) Collect(ch chan<- prometheus.Metric) {
	metric := prometheus.MustNewConstMetric(c.buildInfoDesc, prometheus.GaugeValue, float64(1))
	ch <- metric
	//ch <- prometheus.MustNewConstMetric(c.buildInfoDesc, prometheus.GaugeValue, float64(runtime.NumGoroutine()))
}
