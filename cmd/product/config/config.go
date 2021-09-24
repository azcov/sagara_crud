package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/azcov/sagara_crud/pkg/util"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	App struct {
		Name            string        `mapstructure:"name"`
		PodIP           string        `mapstructure:"pod_ip"`
		Domain          string        `mapstructure:"domain"`
		Environment     string        `mapstructure:"environtment"`
		ShutdownTimeout time.Duration `mapstructure:"shoutdown_timeout"`
	} `mapstructure:"api"`
	HTTP struct {
		API struct {
			Host         string        `mapstructure:"host"`
			Port         int           `mapstructure:"port"`
			ReadTimeout  time.Duration `mapstructure:"read_timeout"`
			WriteTimeout time.Duration `mapstructure:"write_timeout"`
			IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
		} `mapstructure:"api"`
	} `mapstructure:"http"`
	Database struct {
		Pg struct {
			Host                  string        `mapstructure:"host"`
			Port                  string        `mapstructure:"port"`
			Dbname                string        `mapstructure:"dbname"`
			User                  string        `mapstructure:"user"`
			Password              string        `mapstructure:"password"`
			Sslmode               string        `mapstructure:"sslmode"`
			MaxOpenConnection     int           `mapstructure:"max_open_connection"`
			MaxIdleConnection     int           `mapstructure:"max_idle_connection"`
			MaxConnectionLifetime time.Duration `mapstructure:"max_connection_lifetime"`
		} `mapstructure:"pg"`
	} `mapstructure:"database"`
	Auth struct {
		TokenType          string `mapstructure:"token_type"`
		AccessTokenSecret  string `mapstructure:"access_token_secret"`
		RefreshTokenSecret string `mapstructure:"refresh_token_secret"`
	} `mapstructure:"auth"`
	Service struct {
		Auth struct {
			Address string `mapstructure:"address"`
			Port    string `mapstructure:"port"`
		} `mapstructure:"auth"`
	} `mapstructure:"service"`
}

// setupMainConfig loads app config to viper
func GetConfigJSON(logger *zap.SugaredLogger) *Config {
	var c Config
	logger.Info("Executing config")
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "local"
	}
	pathConfig := fmt.Sprintf("cmd/product/config/app/%s.json", appEnv)

	logger.Infof("config path = %s", pathConfig)

	if util.IsFileorDirExist(pathConfig) {
		logger.Infof("CONFIG: APP_ENV = %s , %s file is found, now assigning it with default config", appEnv, pathConfig)
		viper.SetConfigFile(pathConfig)
		err := viper.ReadInConfig()
		if err != nil {
			logger.Info("err: ", err)
		}
	} else {
		logger.Fatal("Config is required")
	}

	viper.SetEnvPrefix(`app`)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	err := viper.Unmarshal(&c)
	if err != nil {
		logger.Fatal("err: ", err)
	}

	if !util.IsFileorDirExist(pathConfig) {
		// open a goroutine to watch remote changes forever
		go func() {
			for {
				time.Sleep(time.Second * 5)

				err := viper.WatchRemoteConfig()
				if err != nil {
					logger.Errorf("unable to read remote config: %v", err)
					continue
				}

				// unmarshal new config into our runtime config struct. you can also use channel
				// to implement a signal to notify the system of the changes
				viper.Unmarshal(&c)
			}
		}()
	}
	zap.S().Infof("%+v", c)

	return &c
}
