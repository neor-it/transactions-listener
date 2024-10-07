package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"strings"
)

// LoadConfig loads the configuration from the environment variables.§
func LoadConfig() (*TxLoaderConfig, error) {
	// Load the environment variables from the .env file.
	err := godotenv.Load()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load the environment variables from the .env file")
	}

	// Load the configuration from the environment variables using the envconfig package.
	var cfg TxLoaderConfig
	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load the configuration from the environment variables")
	}

	// Trim spaces from the chain IDs array.
	for i, id := range cfg.ChainIDs {
		cfg.ChainIDs[i] = strings.TrimSpace(id)
	}

	return &cfg, nil
}
