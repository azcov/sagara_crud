package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Adapter is http server app adapter
type Adapter struct {
	name       string
	httpServer *echo.Echo
	logger     *zap.SugaredLogger
}

// NewAdapter provides new primary HTTP adapter
func NewAdapter(name string, httpServer *echo.Echo, logger *zap.SugaredLogger) *Adapter {
	return &Adapter{
		name:       name,
		httpServer: httpServer,
		logger:     logger,
	}
}

// Start start http application adapter
func (adapter *Adapter) Start(ctx context.Context) error {
	adapter.logger.Infof("start echo adapter %s in %s", adapter.name, adapter.httpServer.Server.Addr)
	if err := adapter.httpServer.Start(adapter.httpServer.Server.Addr); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Stop stops http application adapter
func (adapter *Adapter) Stop(ctx context.Context) error {
	return adapter.httpServer.Shutdown(ctx)
}

func (adapter *Adapter) GetName() string {
	return adapter.name
}
