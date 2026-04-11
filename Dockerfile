FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ethiocal .


FROM alpine:3.21
WORKDIR /opt
COPY --from=builder /app/ethiocal /opt/ethiocal/
EXPOSE 8080
CMD ["./opt/ethiocal"]
