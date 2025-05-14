# OpenTelemetry Collector - Exemplo de Configuração

Este projeto demonstra diferentes configurações do OpenTelemetry Collector para coleta de logs, traces e métricas.

## Estrutura do Projeto

```
.
├── otel-logs/                  # Configuração básica para coleta de logs
│   ├── docker-compose.yml      # Serviços: OTEL Collector, Loki, Grafana
│   └── config/
│       └── otel-collector-config.yml  # Configuração do Collector
│
└── otel-logs-trace-metrics/    # Configuração avançada com traces e métricas
    ├── docker-compose.yml      # Serviços: OTEL Collector, Prometheus, Tempo, Grafana
    ├── config/
    │   ├── otel-collector-config.yaml # Configuração completa do Collector
    │   ├── prometheus.yaml     # Configuração do Prometheus
    │   └── tempo.yaml          # Configuração do Tempo
    └── otel-metric-trace-generator/  # Aplicação exemplo em Go
        └── main.go             # Gera métricas e traces de exemplo
```

## Pré-requisitos

- Docker e Docker Compose instalados
- Portas disponíveis: 4317, 4318, 9090, 3000, 3100, 14268

## Como Executar

### Configuração Básica (Logs)

```bash
cd otel-logs
docker-compose up -d
```

Acesse o Grafana em http://localhost:3000

### Configuração Completa (Logs, Traces e Métricas)

```bash
cd otel-logs-trace-metrics
docker-compose up -d --build
```

Acesse:
- Grafana: http://localhost:3000
- Prometheus: http://localhost:9090
- Tempo: http://localhost:3100

## Configurações

### Collector (otel-collector-config.yaml)

O arquivo de configuração define:
- Receivers: OTLP, Prometheus
- Processors: Batch
- Exporters: Loki (logs), Prometheus (métricas), Tempo (traces)

### Aplicação Exemplo

O gerador (`otel-metric-trace-generator`) demonstra:
- Instrumentação com OpenTelemetry
- Exportação de métricas e traces
- Configuração via variáveis de ambiente

## Links Úteis

- [Documentação OpenTelemetry](https://opentelemetry.io/docs/)
- [Grafana Loki](https://grafana.com/docs/loki/latest/)
- [Grafana Tempo](https://grafana.com/docs/tempo/latest/)
