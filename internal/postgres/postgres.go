package postgres

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq" // postgres driver
	"go.uber.org/zap"
)

type PostgresConnection struct {
	Host                  string        `mapstructure:"host"`
	Port                  string        `mapstructure:"port"`
	DbName                string        `mapstructure:"dbname"`
	User                  string        `mapstructure:"user"`
	Password              string        `mapstructure:"password"`
	SslMode               string        `mapstructure:"sslmode"`
	MaxOpenConnection     int           `mapstructure:"max_open_connection"`
	MaxIdleConnection     int           `mapstructure:"max_idle_connection"`
	MaxConnectionLifetime time.Duration `mapstructure:"max_connection_lifetime"`
}

// CreatePGConnection return db connection instance
func CreatePGConnection(opts PostgresConnection) (*sql.DB, error) {
	port, err := strconv.Atoi(opts.Port)
	if err != nil {
		zap.S().Fatal("Invalid port number : ", opts.Port)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=%s",
		opts.Host, port, opts.User, opts.Password, opts.DbName, opts.SslMode)
	zap.S().Info("pqsl info = ", psqlInfo)
	// if util.IsProductionEnv() {
	// 	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	// 		"password=%s dbname=%s sslrootcert=./rds-combined-ca-bundle.pem sslmode=%s",
	// 		opts.Host, port, opts.User, opts.Password, opts.DbName, opts.SslMode)
	// }

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		zap.S().Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		zap.S().Fatal(err)
	}

	// Setting database connection config
	db.SetMaxOpenConns(opts.MaxOpenConnection)
	db.SetMaxIdleConns(opts.MaxIdleConnection)
	db.SetConnMaxLifetime(opts.MaxConnectionLifetime)

	zap.S().Info("Connected to PG DB Server: ", opts.Host, " at port:", opts.Port, " successfully!")

	return db, nil
}
