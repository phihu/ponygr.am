package config

import (
	"net/url"
	"time"

	"github.com/vrischmann/envconfig"

	"github.com/pkg/errors"
)

// Config represents the configuration of the ponysrv app
type Config struct {
	Env   string `envconfig:"default=dev"`
	Debug bool

	HTTP struct {
		Addr    string `envconfig:"default=:8080"`
		BaseURL string `envconfig:"default=http://localhost:8080"`
		baseURL *url.URL

		Secure bool `envconfig:"default=false"`

		ReadTimeout  time.Duration `envconfig:"default=5s"`
		WriteTimeout time.Duration `envconfig:"default=10s"`
		IdleTimeout  time.Duration `envconfig:"default=2m"`

		Client struct {
			Timeout time.Duration `envconfig:"default=10s"`
		}

		Proxy bool

		CORSOrigins []string
		CSPHeader   string `envconfig:"default=default-src 'self'"`
	}

	DB struct {
		Debug bool

		DSN             string        `envconfig:"default=root:root@tcp(127.0.0.1)/ponygram?charset=utf8mb4&parseTime=true&time_zone=UTC"`
		ConnMaxLifetime time.Duration `envconfig:"default=30m"`

		CACert string
	}

	Session struct {
		Debug   bool
		Name    string        `envconfig:"default=hs"`
		Timeout time.Duration `envconfig:"default=720h"`
	}
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.InitWithOptions(cfg, envconfig.Options{
		AllOptional:     true,
		AllowUnexported: true,
	})
	if err == nil {
		err = cfg.Check()
	}
	if err != nil {
		return nil, errors.Wrap(err, "error initializing Config")
	}
	err = cfg.Parse()
	if err != nil {
		return nil, errors.Wrap(err, "error parsing Config")
	}
	return cfg, nil
}

func (c *Config) Check() error {
	if c.DB.DSN == "" {
		return errors.New("missing DB DSN")
	}

	return nil
}

func (c *Config) Parse() error {
	var err error
	c.HTTP.baseURL, err = url.Parse(c.HTTP.BaseURL)
	if err != nil {
		return errors.Wrap(err, "error parsing HTTP Base URL")
	}

	return nil
}
