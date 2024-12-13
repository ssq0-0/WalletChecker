FROM golang:1.22 AS builder

WORKDIR /walletchecker

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o walletchecker ./core/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /walletchecker/walletchecker /app/walletchecker

COPY --from=builder /walletchecker/config/config.json /app/config/config.json

RUN chmod +x /app/walletchecker

VOLUME ["/app/account"]

ENTRYPOINT ["/app/walletchecker"]
