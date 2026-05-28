FROM golang:1.25.0-bookworm AS builder

WORKDIR /src


COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/monitoring-service ./cmd

FROM debian:bookworm-slim

ENV TZ=Asia/Jakarta
ENV PORT=8080

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates tzdata \
    && rm -rf /var/lib/apt/lists/* \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && useradd --system --create-home --shell /usr/sbin/nologin appuser

WORKDIR /app
COPY --from=builder /out/monitoring-service /app/monitoring-service

RUN chown -R appuser:appuser /app
USER appuser

EXPOSE 8080

CMD ["/app/monitoring-service"]