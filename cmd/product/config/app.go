package config

import (
	"database/sql"
	"os"

	"github.com/azcov/sagara_crud/internal/postgres"
	"github.com/azcov/sagara_crud/pkg/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ConnectToPGServer is a function to init PostgreSQL connection
func ConnectToPGServer(cfg *Config) (*sql.DB, error) {
	if util.IsProductionEnv() && cfg.Database.Pg.Password != "" {
		zap.S().Fatal("postgres password can not be empty!")
	}

	dbpg, err := postgres.CreatePGConnection(postgres.PostgresConnection{
		Host:                  cfg.Database.Pg.Host,
		Port:                  cfg.Database.Pg.Port,
		DbName:                cfg.Database.Pg.Dbname,
		User:                  cfg.Database.Pg.User,
		Password:              cfg.Database.Pg.Password,
		SslMode:               cfg.Database.Pg.Sslmode,
		MaxOpenConnection:     cfg.Database.Pg.MaxOpenConnection,
		MaxIdleConnection:     cfg.Database.Pg.MaxIdleConnection,
		MaxConnectionLifetime: cfg.Database.Pg.MaxConnectionLifetime,
	})

	if err != nil {
		os.Exit(1)
	}

	return dbpg, err
}

func SetupLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true

	// using json format if app_env is production or development
	if os.Getenv("APP_ENV") == "production" || os.Getenv("APP_ENV") == "development" {
		config = zap.NewProductionConfig()
		config.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		config.DisableStacktrace = false
	}

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000000Z")
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
	return logger
}
