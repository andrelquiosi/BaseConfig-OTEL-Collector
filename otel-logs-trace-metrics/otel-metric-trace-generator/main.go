package main

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"

	"otel_slog_example/internal/observability"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {

	hostname, err := os.Hostname()
	if err != nil {
		slog.Error("Failed to get hostname", "error", err)
		return
	}
	// Logger
	logFile, err := os.OpenFile("/log/otel-metric-trace-generator.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("Failed to open log file", "error", err)
		return
	}
	defer logFile.Close()

	multi := io.MultiWriter(os.Stdout, logFile)
	logHandler := slog.NewJSONHandler(multi, nil)
	logger := slog.New(logHandler)

	// Adicionar informações adicionais (como nome da aplicação, versão, e host)
	logger = logger.With(
		slog.String("application", "otel-metric-trace-generator"),
		slog.String("host", hostname),
	)

	// Configurar o logger com essas informações adicionais
	slog.SetDefault(logger)

	shutdown := observability.InitTracer()
	defer shutdown(context.Background())

	// Setup monitor
	// monitor := usecase.NewGatewayHealthChecker()
	// go monitor.MonitorGateway(context.Background())

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tracer := observability.GetTracer()
		ctx, span := tracer.Start(ctx, "HandleRoot")
		defer span.End()

		observability.LoggerWithTrace(ctx).Info("Handling root route")
		w.Write([]byte("Hello, Tracing Logger!"))
	}))

	handler := observability.LogWithTraceMiddleware(otelhttp.NewHandler(mux, "example-server"))
	http.ListenAndServe(":8080", handler)
}
