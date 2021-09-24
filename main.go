package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/azcov/sagara_crud/cmd/auth"
	"github.com/azcov/sagara_crud/cmd/product"
	"github.com/azcov/sagara_crud/internal/application"
	"github.com/azcov/sagara_crud/internal/logger"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logger := logger.NewLogger()

	app := application.New(logger)
	app.AddAdapters(auth.AuthAdapter()...)
	app.AddAdapters(product.ProductAdapter()...)

	app.WithShutdownTimeout(time.Second * 10)
	app.Run(ctx)
}
