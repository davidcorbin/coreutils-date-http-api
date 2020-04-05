FROM golang:1.12-alpine AS build_base
RUN apk add --no-cache git
WORKDIR /go/src/github.com/davidcorbin/coreutils-date-http-api
COPY . .
RUN go get -d -v ./...
RUN go build -o /out/coreutils-date-http-api cmd/server/main.go


FROM alpine:3.9

COPY --from=build_base /out/coreutils-date-http-api /app/coreutils-date-http-api
EXPOSE 8080
CMD ["/app/coreutils-date-http-api"]
