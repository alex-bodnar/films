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

const (
	// DefaultPath - default path for config.
	DefaultPath = "./cmd/config.yaml"
)

type (
	// Config defines the properties of the application configuration.
	Config struct {
		Logger   log.Config `yaml:"logger"`
		Storage  Storage    `yaml:"storage"`
		Delivery Delivery   `yaml:"delivery"`
		Extra    Extra      `yaml:"extra"`
	}

	// Storage defines database engines configuration
	Storage struct {
		Postgres database.Config `yaml:"postgres"`
		Redis    redis.Config    `yaml:"redis"`
	}

	// Delivery defines API server configuration.
	Delivery struct {
		HTTPServer HTTPServer `yaml:"http-server"`
	}

	// HTTPServer defines HTTP section of the API server configuration.
	HTTPServer struct {
		LogRequests        bool          `yaml:"log-requests"`
		ListenAddress      string        `yaml:"listen-address"`
		ReadTimeout        time.Duration `yaml:"read-timeout"`
		WriteTimeout       time.Duration `yaml:"write-timeout"`
		BodySizeLimitBytes int           `yaml:"body-size-limit"`
		GracefulTimeout    int           `yaml:"graceful-timeout"`
	}

	// Extra defines business configuration
	Extra struct {
		RedisCache RedisCache `yaml:"redis-cache"`
		LocalCache LocalCache `yaml:"local-cache"`
	}

	// RedisCache defines redis cache configuration.
	RedisCache struct {
		TimeLive time.Duration `yaml:"time-live"`
	}

	// LocalCache defines redis cache configuration.
	LocalCache struct {
		TimeLive        time.Duration `yaml:"time-live"`
		NumberOfRecords int           `yaml:"number-of-records"`
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
