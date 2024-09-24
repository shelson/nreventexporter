// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nreventexporter // import "github.com/shelson/nreventexporter"
import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Config defines configuration for OTLP/HTTP exporter.
type Config struct {
	confighttp.ClientConfig      `mapstructure:",squash"`  // squash ensures fields are correctly decoded in embedded struct.
	RetryConfig                  configretry.BackOffConfig `mapstructure:"retry_on_failure"`
	exporterhelper.QueueSettings `mapstructure:"sending_queue"`

	// The URL to send metrics to. If omitted the Endpoint + "/v1/metrics" will be used.
	MetricsEndpoint string `mapstructure:"metrics_endpoint"`

	// API key to use when sending data to the New Relic backend.
	APIKey string `mapstructure:"api_key"`
}

var _ component.Config = (*Config)(nil)

// Validate checks if the exporter configuration is valid
func (cfg *Config) Validate() error {
	if cfg.Endpoint == "" {
		return errors.New("at least one endpoint must be specified")
	}
	return nil
}
