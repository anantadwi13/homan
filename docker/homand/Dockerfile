FROM golang:1.17 AS builder

WORKDIR /go/src/service
COPY go.* ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go test ./internal/homand/...
RUN GOOS=linux CGO_ENABLED=0 go build -o service ./cmd/homand/

FROM alpine:3.14
WORKDIR /root
COPY --from=builder /go/src/service/service .

EXPOSE 80

ENTRYPOINT ["./service"]