package article

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/azcov/sagara_crud/docs"
	"github.com/azcov/sagara_crud/internal/application"
	"github.com/azcov/sagara_crud/internal/authentication"
	"github.com/azcov/sagara_crud/internal/logger"
	"google.golang.org/grpc"

	articleHttp "github.com/azcov/sagara_crud/cmd/article/app/interfaces/http"
	articleRepo "github.com/azcov/sagara_crud/cmd/article/app/repository/article"
	articleUsecase "github.com/azcov/sagara_crud/cmd/article/app/usecases"
	articleConfig "github.com/azcov/sagara_crud/cmd/article/config"
	authProto "github.com/azcov/sagara_crud/cmd/auth/proto"
	echoAdapter "github.com/azcov/sagara_crud/internal/http/echo"
	appMiddleware "github.com/azcov/sagara_crud/internal/http/echo/middleware"
)

// @title Mufid Article API
// @version 1.0
// @BasePath /

func ArticleAdapter() []application.Adapter {
	logger := logger.NewLogger()

	cfg := articleConfig.GetConfigJSON(logger)

	pgDB, err := articleConfig.ConnectToPGServer(cfg)
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
	articleRepository := articleRepo.NewRepository(pgDB)
	articleUsecase := articleUsecase.NewUsecase(cfg, articleRepository)
	customMiddleware := appMiddleware.NewMiddleware(authenticator, logger, authGrpcClient)

	router := articleHttp.NewRouter(authenticator, articleUsecase, customMiddleware)

	router.Server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.HTTP.API.Host, cfg.HTTP.API.Port),
		ReadTimeout:  cfg.HTTP.API.ReadTimeout,
		WriteTimeout: cfg.HTTP.API.WriteTimeout,
		IdleTimeout:  cfg.HTTP.API.IdleTimeout,
		Handler:      router,
	}
	return []application.Adapter{
		echoAdapter.NewAdapter(
			"ArticleHTTP",
			router,
			logger,
		),
	}
}
