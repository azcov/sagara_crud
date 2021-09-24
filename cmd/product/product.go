package product

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/azcov/sagara_crud/docs"
	"github.com/azcov/sagara_crud/internal/application"
	"github.com/azcov/sagara_crud/internal/authentication"
	"github.com/azcov/sagara_crud/internal/logger"
	"google.golang.org/grpc"

	authProto "github.com/azcov/sagara_crud/cmd/auth/proto"
	productHttp "github.com/azcov/sagara_crud/cmd/product/app/interfaces/http"
	productRepo "github.com/azcov/sagara_crud/cmd/product/app/repository/product"
	productUsecase "github.com/azcov/sagara_crud/cmd/product/app/usecases"
	productConfig "github.com/azcov/sagara_crud/cmd/product/config"
	echoAdapter "github.com/azcov/sagara_crud/internal/http/echo"
	appMiddleware "github.com/azcov/sagara_crud/internal/http/echo/middleware"
)

// @name Mufid Product API
// @version 1.0
// @BasePath /

func ProductAdapter() []application.Adapter {
	logger := logger.NewLogger()

	cfg := productConfig.GetConfigJSON(logger)

	pgDB, err := productConfig.ConnectToPGServer(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	// auth service
	transportOption := grpc.WithInsecure()
	authAddress := fmt.Sprintf("%s:%s", cfg.Service.Auth.Address, cfg.Service.Auth.Port)
	connAuth, err := grpc.Dial(authAddress, transportOption)
	if err != nil {
		panic(err)
	}
	authGrpcClient := authProto.NewAuthServiceClient(connAuth)

	authenticator := authentication.NewSecretAuthenticator(
		cfg.Auth.TokenType,
		cfg.App.Name,
		cfg.Auth.AccessTokenSecret,
		cfg.Auth.RefreshTokenSecret,
		time.Second,
		time.Second)
	productRepository := productRepo.NewRepository(pgDB)
	productUsecase := productUsecase.NewUsecase(cfg, productRepository)
	customMiddleware := appMiddleware.NewMiddleware(authenticator, logger, authGrpcClient)

	router := productHttp.NewRouter(authenticator, productUsecase, customMiddleware)

	router.Server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HTTP.API.Host, cfg.HTTP.API.Port),
		ReadTimeout:  cfg.HTTP.API.ReadTimeout,
		WriteTimeout: cfg.HTTP.API.WriteTimeout,
		IdleTimeout:  cfg.HTTP.API.IdleTimeout,
		Handler:      router,
	}
	return []application.Adapter{
		echoAdapter.NewAdapter(
			"ProductHTTP",
			router,
			logger,
		),
	}
}
