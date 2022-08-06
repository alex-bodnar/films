package config

import (
	"fmt"
	"time"

	"films-api/pkg/config"
	"films-api/pkg/database"
	"films-api/pkg/errs"
	"films-api/pkg/log"
	"films-api/pkg/redis"
)

//go:generate go-validator

const (
	// DefaultPath - default path for config.
	DefaultPath = "./cmd/config.yaml"
)

type (
	// Config defines the properties of the application configuration.
	Config struct {
		Logger   log.Config `yaml:"logger" valid:"check,deep"`
		Storage  Storage    `yaml:"storage" valid:"check,deep"`
		Delivery Delivery   `yaml:"delivery" valid:"check,deep"`
		Extra    Extra      `yaml:"extra" valid:"check,deep"`
	}

	// Storage defines database engines configuration
	Storage struct {
		Postgres database.Config `yaml:"postgres" valid:"check,deep"`
		Redis    redis.Config    `yaml:"redis" valid:"check,deep"`
	}

	// Delivery defines API server configuration.
	Delivery struct {
		HTTPServer HTTPServer `yaml:"http-server" valid:"check,deep"`
	}

	// HTTPServer defines HTTP section of the API server configuration.
	HTTPServer struct {
		LogRequests        bool          `yaml:"log-requests"`
		ListenAddress      string        `yaml:"listen-address" valid:"required"`
		ReadTimeout        time.Duration `yaml:"read-timeout" valid:"required"`
		WriteTimeout       time.Duration `yaml:"write-timeout" valid:"required"`
		BodySizeLimitBytes int           `yaml:"body-size-limit" valid:"required"`
		GracefulTimeout    int           `yaml:"graceful-timeout" valid:"required"`
	}

	// Extra defines business configuration
	Extra struct {
		RedisCache RedisCache `yaml:"redis-cache" valid:"check,deep"`
		LocalCache LocalCache `yaml:"local-cache" valid:"check,deep"`
	}

	// RedisCache defines redis cache configuration.
	RedisCache struct {
		TimeLive time.Duration `yaml:"time-live" valid:"required"`
	}

	// LocalCache defines redis cache configuration.
	LocalCache struct {
		TimeLive        time.Duration `yaml:"time-live" valid:"required"`
		NumberOfRecords int           `yaml:"number-of-records" valid:"required"`
	}
)

// New loads and validates all configuration data, returns filled Cfg - configuration data model.
func New(appName, cfgFilePath string) (*Config, error) {
	cfg := new(Config)

	if cfgErr := cfg.loadFromFile(cfgFilePath); cfgErr != nil {
		return nil, fmt.Errorf("config loader: %s", cfgErr)
	}

	return cfg.valid()
}

// loadFromFile loads configuration from file.
func (c *Config) loadFromFile(configPath string) error {
	if err := config.LoadFromFile(configPath, c); err != nil {
		return err
	}

	return nil
}

// valid validates configuration data.
func (c *Config) valid() (*Config, error) {
	if errorsList := c.Validate(); len(errorsList) != 0 {
		return nil, &errs.FieldsValidation{Errors: errorsList}
	}

	return c, nil
}
