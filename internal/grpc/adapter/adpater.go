package adapter

import (
	"context"
	"net"

	"go.uber.org/zap"
	grpcmain "google.golang.org/grpc"
	grpchealth "google.golang.org/grpc/health"
	grpchealthproto "google.golang.org/grpc/health/grpc_health_v1"
)

// Adapter is grpc server app adapter
type Adapter struct {
	name         string
	address      string
	server       *grpcmain.Server
	healthServer *grpchealth.Server
	logger       *zap.SugaredLogger
}

// NewAdapter provides new primary adapter
func NewAdapter(name, address string, server *grpcmain.Server, logger *zap.SugaredLogger) *Adapter {
	return &Adapter{
		name:         name,
		address:      address,
		server:       server,
		healthServer: grpchealth.NewServer(),
		logger:       logger,
	}
}

// Start start grpc application adapter
func (adapter *Adapter) Start(ctx context.Context) error {
	grpchealthproto.RegisterHealthServer(adapter.server, adapter.healthServer)
	adapter.logger.Infof("start adapter %s in %s", adapter.name, adapter.address)

	lis, err := net.Listen("tcp", adapter.address)
	if err != nil {
		return err
	}

	adapter.healthServer.SetServingStatus(adapter.name, grpchealthproto.HealthCheckResponse_SERVING)

	return adapter.server.Serve(lis)
}

// Stop stops grpc application adapter
func (adapter *Adapter) Stop(ctx context.Context) error {
	adapter.healthServer.SetServingStatus(adapter.name, grpchealthproto.HealthCheckResponse_NOT_SERVING)

	adapter.server.GracefulStop()

	return nil
}

func (adapter *Adapter) GetName() string {
	return adapter.name
}
