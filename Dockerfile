FROM golang:alpine AS builder

WORKDIR /go/src/github.com/mss-boot-io/service-grpc

COPY . .

RUN CGO_ENABLED=0 go build -o application main.go

FROM alpine

LABEL authors="lwnmengjing"

WORKDIR /app

COPY --from=builder /go/src/github.com/mss-boot-io/service-grpc/application /app/application

ENTRYPOINT [ "/app/application" ]