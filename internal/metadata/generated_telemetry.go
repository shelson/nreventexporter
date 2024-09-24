// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"errors"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
)

// Deprecated: [v0.108.0] use LeveledMeter instead.
func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/shelson/nreventexporter")
}

func LeveledMeter(settings component.TelemetrySettings, level configtelemetry.Level) metric.Meter {
	return settings.LeveledMeterProvider(level).Meter("github.com/shelson/nreventexporter")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/shelson/nreventexporter")
}

// TelemetryBuilder provides an interface for components to report telemetry
// as defined in metadata and user config.
type TelemetryBuilder struct {
	meter                    metric.Meter
	ExporterRequestsBytes    metric.Int64Counter
	ExporterRequestsDuration metric.Int64Counter
	ExporterRequestsRecords  metric.Int64Counter
	ExporterRequestsSent     metric.Int64Counter
	meters                   map[configtelemetry.Level]metric.Meter
}

// TelemetryBuilderOption applies changes to default builder.
type TelemetryBuilderOption interface {
	apply(*TelemetryBuilder)
}

type telemetryBuilderOptionFunc func(mb *TelemetryBuilder)

func (tbof telemetryBuilderOptionFunc) apply(mb *TelemetryBuilder) {
	tbof(mb)
}

// NewTelemetryBuilder provides a struct with methods to update all internal telemetry
// for a component
func NewTelemetryBuilder(settings component.TelemetrySettings, options ...TelemetryBuilderOption) (*TelemetryBuilder, error) {
	builder := TelemetryBuilder{meters: map[configtelemetry.Level]metric.Meter{}}
	for _, op := range options {
		op.apply(&builder)
	}
	builder.meters[configtelemetry.LevelBasic] = LeveledMeter(settings, configtelemetry.LevelBasic)
	var err, errs error
	builder.ExporterRequestsBytes, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_exporter_requests_bytes",
		metric.WithDescription("Total size of requests (in bytes)"),
		metric.WithUnit("By"),
	)
	errs = errors.Join(errs, err)
	builder.ExporterRequestsDuration, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_exporter_requests_duration",
		metric.WithDescription("Duration of HTTP requests (in milliseconds)"),
		metric.WithUnit("ms"),
	)
	errs = errors.Join(errs, err)
	builder.ExporterRequestsRecords, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_exporter_requests_records",
		metric.WithDescription("Total size of requests (in number of records)"),
		metric.WithUnit("{records}"),
	)
	errs = errors.Join(errs, err)
	builder.ExporterRequestsSent, err = builder.meters[configtelemetry.LevelBasic].Int64Counter(
		"otelcol_exporter_requests_sent",
		metric.WithDescription("Number of requests"),
		metric.WithUnit("1"),
	)
	errs = errors.Join(errs, err)
	return &builder, errs
}
