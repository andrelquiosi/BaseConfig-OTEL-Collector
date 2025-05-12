package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

var tp *sdktrace.TracerProvider

func InitTracer() func(context.Context) error {
	// Exemplo com exportador em memória (tracetest) apenas para testes locais
	exp, _ := otlptracehttp.New(context.Background(), otlptracehttp.WithInsecure())

	tp = sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("otel-slog-example"),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp.Shutdown
}

func GetTracer() trace.Tracer {
	return tp.Tracer("otel-slog-example")
}
