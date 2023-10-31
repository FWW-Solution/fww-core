package tracing

import (
	"context"
	"fmt"
	"fww-core/internal/config"
	"strings"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	cfg *config.Config
)

func Init(ctx context.Context, service string, cfg *config.Config) trace.Tracer { // set exporter
	exporter, err := createOtelExporter(cfg.OpenTelemetry.ExporterType)
	if err != nil {
		panic(err)
	}
	// set trace provider
	traceProvider := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(newResource(service)))
	tracer := traceProvider.Tracer(service)
	return tracer
}

func createOtelExporter(exporterType string) (sdktrace.SpanExporter, error) {
	// set default exporter type
	if exporterType == "" {
		exporterType = "otlphttp"
	}
	switch exporterType {
	case "otlphttp":
		var opts []otlptracehttp.Option
		if !withSecure(cfg) {
			opts = []otlptracehttp.Option{otlptracehttp.WithInsecure()}
		}
		return otlptrace.New(context.Background(), otlptracehttp.NewClient(opts...))
	case "otlpgrpc":
		return otlptrace.New(context.Background(), otlptracegrpc.NewClient())
	default:
		return nil, fmt.Errorf("unknown exporter type %s", exporterType)
	}
}

// withSecure instructs the client to use HTTPS scheme, instead of hotrod's desired default HTTP
func withSecure(cfg *config.Config) bool {
	return strings.HasPrefix(cfg.OpenTelemetry.OtelExporterOLTPEndpoint, "https://") ||
		cfg.OpenTelemetry.OtelExporterOTLPInsecure
}

func newResource(service string) *resource.Resource {
	return resource.NewWithAttributes(semconv.SchemaURL,
		semconv.ServiceName(service),
	)
}
