package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	Server   ServerConfig
	Log      LogConfig
	Crawler  CrawlerConfig
	Database DatabaseConfig
	Cache    RedisConfig // REFACTOR: rename to CacheConfig where will be more drivers
}

type ServerConfig struct {
	Host string `env:"SERVER_HOST" envDefault:"localhost"`
	Port string `env:"SERVER_PORT" envDefault:"8000"`

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
	Size int `env:"CRAWLER_GORS_COUNT" envDefault:"4"`
}

type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST"`
	Port     string `env:"DATABASE_PORT"`
	Name     string `env:"DATABASE_NAME"`
	User     string `env:"DATABASE_USER"`
	Password string `env:"DATABASE_PASSWORD"`
	Database string `env:"DATABASE_DB_NAME"`
}

type RedisConfig struct {
	Host     string `env:"CACHE_HOST"`
	Port     string `env:"CACHE_PORT"`
	Name     string `env:"CACHE_NAME"`
	User     string `env:"CACHE_USER"`
	Password string `env:"CACHE_PASSWORD"`
	Database int    `env:"CACHE_DB_NAME"`
}

type UnmarshalOptions struct {
	Dev bool
}

func UnmarshalFromEnv(cfg *Config, opts *UnmarshalOptions) error {
	return env.Parse(cfg)
}
