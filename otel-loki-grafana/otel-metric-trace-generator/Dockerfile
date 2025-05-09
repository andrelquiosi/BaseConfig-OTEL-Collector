FROM python:3.11-slim

RUN pip install opentelemetry-sdk \
    opentelemetry-exporter-otlp \
    opentelemetry-instrumentation \
    opentelemetry-api

COPY app.py .

CMD ["python", "app.py"]
