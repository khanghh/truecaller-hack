FROM golang:1.14.2-alpine as builder

WORKDIR /app

RUN apk update && apk add --no-cache git ca-certificates tzdata make && update-ca-certificates

COPY ./cmd/server/main.go .
COPY ./go.mod .
COPY ./Makefile .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make

FROM scratch
WORKDIR /app

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/bin/server ./server

CMD ["./app"]
