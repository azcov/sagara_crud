package auth

import (
	"fmt"
	"net/http"

	_ "github.com/azcov/sagara_crud/docs"
	"github.com/azcov/sagara_crud/internal/application"
	"github.com/azcov/sagara_crud/internal/authentication"
	"github.com/azcov/sagara_crud/internal/logger"
	"github.com/azcov/sagara_crud/pkg/buildinfo"
	"google.golang.org/grpc"

	authGrpc "github.com/azcov/sagara_crud/cmd/auth/app/interfaces/grpc"
	authHttp "github.com/azcov/sagara_crud/cmd/auth/app/interfaces/http"
	authRepo "github.com/azcov/sagara_crud/cmd/auth/app/repository/auth"
	sessionRepo "github.com/azcov/sagara_crud/cmd/auth/app/repository/session"
	authUsecase "github.com/azcov/sagara_crud/cmd/auth/app/usecases"
	authConfig "github.com/azcov/sagara_crud/cmd/auth/config"
	authProto "github.com/azcov/sagara_crud/cmd/auth/proto"
	grpcAdapter "github.com/azcov/sagara_crud/internal/grpc/adapter"
	grpcServer "github.com/azcov/sagara_crud/internal/grpc/server"
	grpcInterceptor "github.com/azcov/sagara_crud/internal/grpc/server/interceptor"
	echoAdapter "github.com/azcov/sagara_crud/internal/http/echo"
)

// @title Mufid Auth API
// @version 1.0
// @BasePath /

func AuthAdapter() []application.Adapter {
	buildinfo.PrintVersionOrContinue()

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	logger := logger.NewLogger()

	cfg := authConfig.GetConfigJSON(logger)

	pgDB, err := authConfig.ConnectToPGServer(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	authentication := authentication.NewSecretAuthenticator(cfg.Auth.TokenType,
		cfg.App.Name,
		cfg.Auth.AccessTokenSecret,
		cfg.Auth.RefreshTokenSecret,
		cfg.Auth.AccessTokenExpiry,
		cfg.Auth.RefreshTokenExpiry)
	authRepository := authRepo.NewRepository(pgDB)
	sessionRepository := sessionRepo.NewRepository(pgDB)
	authUsecase := authUsecase.NewUsecase(cfg, authRepository, authentication)
	grpcAuthServer := authGrpc.NewServer(authUsecase)

	customInterceptor := grpcInterceptor.NewAuthInterceptor(
		authentication,
		sessionRepository,
		logger)

	grpcServer := grpcServer.NewServer(
		grpcServer.ServerConfig{
			ServerMinTime: cfg.HTTP.GRPC.ServerMinTime,
			ServerTime:    cfg.HTTP.GRPC.ServerTime,
			ServerTimeout: cfg.HTTP.GRPC.ServerTimeout,
		},
		logger,
		[]grpc.UnaryServerInterceptor{
			customInterceptor.Unary(),
		},
		nil,
	)

	authProto.RegisterAuthServiceServer(grpcServer, grpcAuthServer)
	router := authHttp.NewRouter(authUsecase)

	router.Server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HTTP.API.Host, cfg.HTTP.API.Port),
		ReadTimeout:  cfg.HTTP.API.ReadTimeout,
		WriteTimeout: cfg.HTTP.API.WriteTimeout,
		IdleTimeout:  cfg.HTTP.API.IdleTimeout,
		Handler:      router,
	}

	return []application.Adapter{
		echoAdapter.NewAdapter("AuthHTTP", router, logger),
		grpcAdapter.NewAdapter("AuthGRPC", fmt.Sprintf("%s:%d", cfg.HTTP.GRPC.Host, cfg.HTTP.GRPC.Port), grpcServer, logger),
	}

	// app := application.New(logger)
	// app.AddAdapters(
	// 	echoAdapter.NewAdapter(
	// 		"AuthHTTP",
	// 		router,
	// 		logger,
	// 	),
	// 	grpcAdapter.NewAdapter(
	// 		"AuthGRPC",
	// 		fmt.Sprintf("%s:%d", cfg.HTTP.GRPC.Host, cfg.HTTP.GRPC.Port),
	// 		grpcServer,
	// 		logger,
	// 	),
	// )

	// app.WithShutdownTimeout(cfg.App.ShutdownTimeout)
	// app.Run(ctx)
}
