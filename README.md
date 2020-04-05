# Coreutils Date HTTP API

![Docker Image CI](https://github.com/davidcorbin/coreutils-date-http-api/workflows/Docker%20Image%20CI/badge.svg)

REST API for managing coreutils date tool.

WARNING: This API requires root access to manage the system date. Additionally, it doesn't require use authentication. Always run behind an authenticated reverse proxy.

## Build

go get -d -v ./...
go build -o coreutils-date-http-api cmd/server/main.go

## Cross compile for ARM64

go get -d -v ./...
go get golang.org/x/sys/unix
env GOOS=linux GOARCH=arm GOARM=5 go build -o coreutils-date-http-api cmd/server/main.go
