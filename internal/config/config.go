package config

import (
	"time"

	"gitlab.com/go-insomnia/config"
	"gitlab.com/go-insomnia/database"
)

const (
	// DefaultPath - default path for config.
	DefaultPath = "./configs/config.yaml"
)

type (
	// Config defines the properties of the bot configuration.
	Config struct {
		Storage Storage `yaml:"storage"`
	}

	// Storage defines databases configuration.
	Storage struct {
		Postgres database.Config `yaml:"postgres"`
		Badger   Badger          `yaml:"badger"`
	}

	Badger struct {
		CipherKey string        `yaml:"cipher-key"`
		KeyTTL    time.Duration `yaml:"key-ttl"`
	}
)

// New returns new configuration.
// It loads config from yaml file, validates it and exit if error is occurred.
func New(filepath string) *Config {
	cfg := new(Config)
	config.LoadFromFileX(filepath, cfg)

	return cfg
}

func (c Config) Validate() error {
	return c.Storage.Postgres.Validate()
}
