package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	Server   ServerConfig
	Log      LogConfig
	Crawler  CrawlerConfig
	Database DatabaseConfig
	Cache    CacheConfig
}

type ServerConfig struct {
	Host  string `env:"HOST"`
	Https HttpsConfig
	Cors  bool `env:"ENABLE_CORS" envDefault:"true"`
}

type HttpsConfig struct {
	Key  string `env:"HTTPS_KEY"`
	Cert string `env:"HTTPS_CERT"`
}

type LogConfig struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
}

type CrawlerConfig struct {
	Size int `env:"CRAWLER_GORS_COUNT" envDefault:"12"`
}

type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST"`
	Port     string `env:"DATABASE_PORT"`
	Name     string `env:"DATABASE_NAME"`
	User     string `env:"DATABASE_USER"`
	Password string `env:"DATABASE_PASSWORD"`
	Database string `env:"DATABASE_DB_NAME"`
}

type CacheConfig struct{}

type UnmarshalOptions struct {
	Dev bool
}

func UnmarshalFromEnv(cfg *Config, opts *UnmarshalOptions) error {
	return env.Parse(cfg)
}
