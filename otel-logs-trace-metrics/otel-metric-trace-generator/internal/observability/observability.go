
package observability

import (
    "context"
    "log/slog"
    "net/http"

    "go.opentelemetry.io/otel/trace"
)

func LoggerWithTrace(ctx context.Context) *slog.Logger {
    logger := slog.Default()
    spanCtx := trace.SpanContextFromContext(ctx)
    if spanCtx.IsValid() {
        logger = logger.With(
            slog.String("trace_id", spanCtx.TraceID().String()),
            slog.String("span_id", spanCtx.SpanID().String()),
        )
    }
    return logger
}

func LogWithTraceMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        LoggerWithTrace(ctx).Info("Incoming request",
            slog.String("method", r.Method),
            slog.String("path", r.URL.Path),
        )
        next.ServeHTTP(w, r)
    })
}
