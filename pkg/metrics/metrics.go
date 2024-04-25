package metrics

import (
	"net/http"

	"github.com/sunsingerus/tbox/pkg/metrics/collectors"
	"github.com/prometheus/client_golang/prometheus"
	prometheusCollectors "github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// SoftwareDescription describes piece of software
type SoftwareDescription struct {
	// Name specifies name of the software. Expected to be human-friendly. Ex.: collector
	Name string
	// Version  specifies version of the software. Ex.: 1.2.3
	Version string
	// Tag specifies git tag of the software sources which it is built from
	Tag string
	// Sha  specifies git sha of the software commit which it is built from
	Sha string
	// Built specifies time when the software is built
	Built string
}

// Exporter is a struct to organize operations of metrics Exporter
type Exporter struct {
	// collectorsRegistry is a registry of metrics collectors
	collectorsRegistry *prometheus.Registry
}

// NewExporter is a constructor
func NewExporter() *Exporter {
	return &Exporter{
		collectorsRegistry: prometheus.NewRegistry(),
	}
}

// RegisterMetricsCollectors registes additional metrics collectors
func (e *Exporter) RegisterMetricsCollectors(collectors ...prometheus.Collector) *Exporter {
	e.collectorsRegistry.MustRegister(collectors...)
	return e
}

// StartMetricsExporterServer starts metrics exporter in background for gRPC metrics
func (e *Exporter) StartMetricsExporterServer(address string, description SoftwareDescription) {
	e.registerStandardMetricsCollectors(description)
	e.startBackgroundServer(address)
}

// registerStandardMetricsCollectors registers standard metrics collectors
func (e *Exporter) registerStandardMetricsCollectors(description SoftwareDescription) {
	// Go standard metrics collector
	e.collectorsRegistry.MustRegister(
		prometheusCollectors.NewGoCollector(
			prometheusCollectors.WithGoCollections(prometheusCollectors.GoRuntimeMetricsCollection),
		),
	)

	// Software build information metrics collector
	e.collectorsRegistry.MustRegister(
		collectors.NewBuildInfoCollector(
			description.Name, description.Version, description.Tag, description.Sha, description.Built,
		),
	)
}

// startBackgroundServer starts metrics exporter as a background go routine
func (e *Exporter) startBackgroundServer(address string) {
	// Expose the registered metrics via HTTP.
	http.Handle("/metrics", promhttp.HandlerFor(
		e.collectorsRegistry,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	))

	go http.ListenAndServe(address, nil)
}
