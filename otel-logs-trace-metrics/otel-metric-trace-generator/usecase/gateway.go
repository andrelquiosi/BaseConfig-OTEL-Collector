
package usecase

import (
    "context"
    "time"

    "otel_slog_example/internal/observability"
)

type GatewayHealthChecker struct{}

func NewGatewayHealthChecker() *GatewayHealthChecker {
    return &GatewayHealthChecker{}
}

func (g *GatewayHealthChecker) MonitorGateway(ctx context.Context) {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()
    for range ticker.C {
        observability.LoggerWithTrace(ctx).Info("Performing gateway check")
    }
}
