FROM golang:1.26 AS builder
RUN apt-get update && apt-get install -y --no-install-recommends gcc libc6-dev && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -tags noui -trimpath -ldflags="-s -w" -o ethiocal .


FROM debian:bookworm-slim
WORKDIR /opt
COPY --from=builder /app/ethiocal .
RUN addgroup --system appgroup && adduser --system --ingroup appgroup appuser
RUN chown appuser:appgroup ethiocal
USER appuser
EXPOSE 8080
CMD ["./ethiocal", "--server"]
