package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config contains all the configuration options used in Instabase Extractor
type Config struct {
	Logger *log.Logger

	// Expected Environment Variable: `IB_VERBOSE_LOG`
	VerboseLog bool `envconfig:"VERBOSE_LOG" default:"false"`
	// Expected Environment Variable: `IB_API_TOKEN`
	APIToken string `envconfig:"API_TOKEN" default:""`
	// Expected Environment Variable: `IB_ROOT_URL`
	RootURL string `envconfig:"ROOT_URL" default:"https://apps.instabase.com/"`
}

// NewAppConfigFromEnv constructs a new Config object and initializes all the fields
// with the value of the environment variables.
func NewAppConfigFromEnv() (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process("IB", cfg); err != nil {
		return cfg, err
	}
	cfg.Logger = log.Default()
	return cfg, nil
}
