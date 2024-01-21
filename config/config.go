package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	App      App
	Mysql    Mysql
	Redis    Redis
	Jwt      Jwt
	Midtrans Midtrans
}

type App struct {
	Name        string
	Version     string
	Port        string
	Environment string
	Timeout     time.Duration
}

type Mysql struct {
	Host              string
	Port              string
	User              string
	Password          string
	DbName            string
	MaxIdleConnection int
	MaxOpenConnection int
}

type Redis struct {
	Host     string
	Port     string
	Password string
}

type Jwt struct {
	SecretKey string
}

type Midtrans struct {
	ServerKey string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(`config/config.yaml`)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
