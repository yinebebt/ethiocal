FROM golang:1.26 AS builder
RUN apt-get update && apt-get install -y gcc libgl1-mesa-dev xorg-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o ethiocal .


FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends libgl1 && rm -rf /var/lib/apt/lists/*
WORKDIR /opt
COPY --from=builder /app/ethiocal .
RUN addgroup --system appgroup && adduser --system --ingroup appgroup appuser
RUN chown appuser:appgroup ethiocal
USER appuser
EXPOSE 8080
CMD ["./ethiocal", "--server"]
