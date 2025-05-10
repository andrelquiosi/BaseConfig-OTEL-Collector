FROM golang:1.24 AS builder

# Set destination for COPY
WORKDIR /app

# Copy the source code
COPY . .
RUN go mod tidy

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -tags timetzdata -o /app/otel-metric-trace-generator ./main.go

# Use a minimal alpine image for the final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/otel-metric-trace-generator .

# Run
CMD ["./otel-metric-trace-generator"]