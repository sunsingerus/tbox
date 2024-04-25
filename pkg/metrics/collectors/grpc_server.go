package collectors

import (
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

type GRPCServerMetricsCollector struct {
	*grpc_prometheus.ServerMetrics
	server *grpc.Server
}

func NewGRPCServerMetricsCollector(namespace, system string) *GRPCServerMetricsCollector {
	return &GRPCServerMetricsCollector{
		ServerMetrics: grpc_prometheus.NewServerMetricsNamed(namespace, system),
	}
}

func (c *GRPCServerMetricsCollector) RegisterServer(server *grpc.Server) {
	c.server = server
	c.InitializeMetrics(c.server)
}

// GetGRPCServerOptions prepares slice of grpc.ServerOption for gRPC metrics collectors
func (c *GRPCServerMetricsCollector) GetGRPCServerOptions() ([]grpc.ServerOption, error) {
	// As we have public key set to parse JWT,
	// we can set up interceptors to perform server-side authorization
	opts := []grpc.ServerOption{
		// Add an interceptor for all unary RPCs.
		grpc.ChainUnaryInterceptor(c.UnaryServerInterceptor()),

		// Add an interceptor for all stream RPCs.
		grpc.ChainStreamInterceptor(c.StreamServerInterceptor()),
	}

	return opts, nil
}
