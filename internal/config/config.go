// Package config handles application configuration.
package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config holds all configuration for the application.
type Config struct {
	// Server configuration
	Port        int    `envconfig:"PORT" default:"8080"`
	Environment string `envconfig:"ENVIRONMENT" default:"development"`
	Version     string `envconfig:"VERSION" default:"0.0.0"`

	// Logging configuration
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`

	// Database configuration (example)
	DatabaseURL string `envconfig:"DATABASE_URL" default:""`

	// External service configuration (example)
	SlackWebhookURL string `envconfig:"SLACK_WEBHOOK_URL" default:""`
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
