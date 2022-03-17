package config

import (
	"crypto/rand"
	"encoding/base64"
	"net/url"
	"time"

	"github.com/phihu/ponygr.am/pkg/session"
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
		Debug               bool
		Name                string        `envconfig:"default=ps"`
		Timeout             time.Duration `envconfig:"default=720h"`
		SignatureKeys       []string
		signatureKeysParsed [][]byte
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
	err = cfg.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "invalid Config")
	}

	return cfg, nil
}

// Check checks the config for validity before parsing
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

	if c.Session.SignatureKeys != nil {
		c.Session.signatureKeysParsed = make([][]byte, len(c.Session.SignatureKeys))
		for i, ks := range c.Session.SignatureKeys {
			c.Session.signatureKeysParsed[i], err = base64.StdEncoding.DecodeString(ks)
			if err != nil {
				return errors.Wrap(err, "error decoding session signature key")
			}
		}
	}
	if c.Debug && c.Session.SignatureKeys == nil {
		c.Session.signatureKeysParsed = [][]byte{
			make([]byte, session.SignatureKeySize),
		}
		read, err := rand.Read(c.Session.signatureKeysParsed[0])
		if err != nil {
			return errors.Wrap(err, "error generating session signature key")
		}
		if read != session.SignatureKeySize {
			return errors.New("not enough bytes read for signature key")
		}
	}

	return nil
}

// Validate returns an error if the config cannot be used
//
// It is run after parsing.
func (c *Config) Validate() error {
	if c.Session.signatureKeysParsed == nil {
		return errors.New("missing session signature keys")
	}
	return nil
}

func (c *Config) HTTPSecure() bool {
	return c.HTTP.Secure
}

func (c *Config) SessionName() string {
	return c.Session.Name
}
func (c *Config) SessionSignatureKeys() [][]byte {
	return c.Session.signatureKeysParsed
}
func (c *Config) SessionTimeout() time.Duration {
	return c.Session.Timeout
}

func (c *Config) DBDSN() string {
	return c.DB.DSN
}
func (c *Config) DBCACert() string {
	return c.DB.CACert
}
func (c *Config) DBConnMaxLifetime() time.Duration {
	return c.DB.ConnMaxLifetime
}
