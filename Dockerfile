FROM golang:1.22.3-alpine3.20 AS BUILDER
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o stress-test main.go
RUN apk update && apk add --no-cache ca-certificates

FROM scratch
WORKDIR /app
COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BUILDER /app/stress-test /app/stress-test
ENTRYPOINT ["./stress-test", "stress"]
