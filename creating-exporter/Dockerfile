FROM golang:alpine AS build

WORKDIR /app
# To create go.mod and go.sum files, run
# `go mod init exporter` and `go mod tidy`
COPY ./exporter/go.mod ./exporter/go.sum ./
RUN go mod download && go mod verify

COPY ./exporter ./
RUN go build exporter.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/exporter ./
CMD ["./exporter"]
