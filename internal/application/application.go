package application

import (
	"context"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
)

// Adapter interface
type Adapter interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	GetName() string
}

// App represents application service
type App struct {
	adapters        []Adapter
	shutdownTimeout time.Duration
	logger          *zap.SugaredLogger
}

// New provides new service application
func New(logger *zap.SugaredLogger) *App {
	return &App{
		shutdownTimeout: 5 * time.Second, // Default shutdown timeout
		logger:          logger,
	}
}

// AddAdapters adds adapters to application service
func (app *App) AddAdapters(adapters ...Adapter) {
	app.adapters = append(app.adapters, adapters...)
}

// WithShutdownTimeout overrides default shutdown timout
func (app *App) WithShutdownTimeout(timeout time.Duration) {
	app.shutdownTimeout = timeout
}

// Run runs the service application
func (app *App) Run(ctx context.Context) {
	for _, adapter := range app.adapters {
		app.logger.Infof("Start Adapter %s", adapter.GetName())
		go func(adapter Adapter) {
			if err := adapter.Start(ctx); err != nil {
				app.logger.Fatalf("adapter start error: %v", adapter.Start(ctx))
				os.Exit(1)
				// adapter.Stop(ctx)
			}
		}(adapter)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	app.stop(ctx)
}

func (app *App) stop(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, app.shutdownTimeout)
	defer cancel()

	app.logger.Info("shutting down...")

	errCh := make(chan error, len(app.adapters))

	for _, adapter := range app.adapters {
		go func(adapter Adapter) {
			errCh <- adapter.Stop(ctxWithTimeout)
		}(adapter)
	}

	for i := 0; i < len(app.adapters); i++ {
		if err := <-errCh; err != nil {
			// calling Goexit terminates that goroutine without returning (previous defers would not run)
			go func(err error) {
				app.logger.Fatalf("shutdown error: %v", err)
				os.Exit(1)
			}(err)
			return
		}
	}

	app.logger.Info("gracefully stopped")
}
